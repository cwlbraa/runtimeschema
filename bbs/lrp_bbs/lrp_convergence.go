package lrp_bbs

import (
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs/prune"
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/shared"
	"github.com/cloudfoundry-incubator/runtime-schema/metric"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/cloudfoundry/storeadapter"
	"github.com/pivotal-golang/lager"
)

const (
	convergeLRPRunsCounter = metric.Counter("ConvergenceLRPRuns")
	convergeLRPDuration    = metric.Duration("ConvergenceLRPDuration")

	lrpsDeletedCounter = metric.Counter("ConvergenceLRPsDeleted")
)

func (bbs *LRPBBS) ConvergeLRPs(pollingInterval time.Duration) {
	logger := bbs.logger.Session("converge-lrps")
	logger.Info("starting-convergence")
	defer logger.Info("finished-convergence")

	convergeLRPRunsCounter.Increment()

	convergeStart := bbs.timeProvider.Now()

	// make sure to get funcy here otherwise the time will be precomputed
	defer func() {
		convergeLRPDuration.Send(bbs.timeProvider.Now().Sub(convergeStart))
	}()

	actualsByProcessGuid, err := bbs.pruneActualsWithMissingCells(logger)
	if err != nil {
		logger.Error("failed-to-fetch-and-prune-actual-lrps", err)
		return
	}

	desiredLRPRoot, err := bbs.store.ListRecursively(shared.DesiredLRPSchemaRoot)
	if err != nil && err != storeadapter.ErrorKeyNotFound {
		logger.Error("failed-to-fetch-desired-lrps", err)
		return
	}

	var malformedDesiredLRPs []string
	desiredLRPsByProcessGuid := map[string]models.DesiredLRP{}

	for _, node := range desiredLRPRoot.ChildNodes {
		var desiredLRP models.DesiredLRP
		err := models.FromJSON(node.Value, &desiredLRP)

		if err != nil {
			logger.Info("pruning-invalid-desired-lrp-json", lager.Data{
				"error":   err.Error(),
				"payload": node.Value,
			})
			malformedDesiredLRPs = append(malformedDesiredLRPs, node.Key)
			continue
		}

		desiredLRPsByProcessGuid[desiredLRP.ProcessGuid] = desiredLRP
		actualLRPsForDesired := actualsByProcessGuid[desiredLRP.ProcessGuid]

		bbs.reconcile(desiredLRP, actualLRPsForDesired, logger)
	}

	actualLRPsToStop := bbs.instancesToStop(desiredLRPsByProcessGuid, actualsByProcessGuid, logger)

	for _, actual := range actualLRPsToStop {
		logger.Info("detected-undesired-instance", lager.Data{
			"process-guid":  actual.ProcessGuid,
			"instance-guid": actual.InstanceGuid,
			"index":         actual.Index,
		})
	}

	lrpsDeletedCounter.Add(uint64(len(malformedDesiredLRPs)))
	bbs.store.Delete(malformedDesiredLRPs...)

	bbs.resendStartAuctions(desiredLRPsByProcessGuid, actualsByProcessGuid, pollingInterval, logger)

	lrpStopInstanceCounter.Add(uint64(len(actualLRPsToStop)))
	err = bbs.RetireActualLRPs(actualLRPsToStop, logger)
	if err != nil {
		logger.Error("failed-to-request-stops", err)
	}
}

func (bbs *LRPBBS) instancesToStop(
	desiredLRPsByProcessGuid map[string]models.DesiredLRP,
	actualsByProcessGuid map[string]models.ActualLRPsByIndex,
	logger lager.Logger,
) []models.ActualLRP {
	var actualsToStop []models.ActualLRP

	for processGuid, actuals := range actualsByProcessGuid {
		if _, found := desiredLRPsByProcessGuid[processGuid]; !found {
			for _, actual := range actuals {
				actualsToStop = append(actualsToStop, actual)
			}
		}
	}

	return actualsToStop
}

func (bbs *LRPBBS) resendStartAuctions(
	desiredLRPsByProcessGuid map[string]models.DesiredLRP,
	actualsByProcessGuid map[string]models.ActualLRPsByIndex,
	pollingInterval time.Duration,
	logger lager.Logger,
) {
	for processGuid, actuals := range actualsByProcessGuid {
		for _, actual := range actuals {
			if actual.State == models.ActualLRPStateUnclaimed && bbs.timeProvider.Now().After(time.Unix(0, actual.Since).Add(pollingInterval)) {
				desiredLRP, found := desiredLRPsByProcessGuid[processGuid]
				if !found {
					logger.Info("failed-to-find-desired-lrp-for-stale-unclaimed-actual-lrp", lager.Data{"actual-lrp": actual})
					continue
				}

				lrpStart := models.LRPStart{
					DesiredLRP: desiredLRP,
					Index:      actual.Index,
				}

				logger.Info("resending-start-auction", lager.Data{"process-guid": processGuid, "index": actual.Index})
				err := bbs.requestLRPStartAuction(lrpStart)
				if err != nil {
					logger.Error("failed-resending-start-auction", err, lager.Data{
						"lrp-start-auction": lrpStart,
					})
				}
			}
		}
	}
}

func (bbs *LRPBBS) pruneActualsWithMissingCells(logger lager.Logger) (map[string]models.ActualLRPsByIndex, error) {
	actualsByProcessGuid := map[string]models.ActualLRPsByIndex{}

	cellRoot, err := bbs.store.ListRecursively(shared.CellSchemaRoot)
	if err == storeadapter.ErrorKeyNotFound {
		cellRoot = storeadapter.StoreNode{}
	} else if err != nil {
		logger.Error("failed-to-get-cells", err)
		return nil, err
	}

	err = prune.Prune(bbs.store, shared.ActualLRPSchemaRoot, func(node storeadapter.StoreNode) (shouldKeep bool) {
		var actual models.ActualLRP
		err := models.FromJSON(node.Value, &actual)
		if err != nil {
			return false
		}

		if actual.State != models.ActualLRPStateUnclaimed {
			if _, ok := cellRoot.Lookup(actual.CellID); !ok {
				logger.Info("detected-actual-with-missing-cell", lager.Data{
					"actual":  actual,
					"cell-id": actual.CellID,
				})
				return false
			}
		}

		actuals, found := actualsByProcessGuid[actual.ProcessGuid]
		if !found {
			actuals = models.ActualLRPsByIndex{}
			actualsByProcessGuid[actual.ProcessGuid] = actuals
		}

		actuals[actual.Index] = actual

		return true
	})

	if err != nil {
		logger.Error("failed-to-prune-actual-lrps", err)
		return nil, err
	}

	return actualsByProcessGuid, nil
}
