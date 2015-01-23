package bbs_test

import (
	. "github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("BBS", func() {
	It("should compile and be able to construct and return each BBS", func() {
		NewBBS(etcdClient, timeProvider, models.NewDefaultRestartCalculator(), logger)
		NewRepBBS(etcdClient, timeProvider, logger)
		NewConvergerBBS(etcdClient, timeProvider, logger)
		NewNsyncBBS(etcdClient, timeProvider, logger)
		NewAuctioneerBBS(etcdClient, timeProvider, logger)
		NewMetricsBBS(etcdClient, timeProvider, logger)
		NewRouteEmitterBBS(etcdClient, timeProvider, logger)
		NewVeritasBBS(etcdClient, timeProvider, logger)
	})
})
