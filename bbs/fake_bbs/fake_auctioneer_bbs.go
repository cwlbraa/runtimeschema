// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
)

type FakeAuctioneerBBS struct {
	CellsStub        func() ([]models.CellPresence, error)
	cellsMutex       sync.RWMutex
	cellsArgsForCall []struct{}
	cellsReturns struct {
		result1 []models.CellPresence
		result2 error
	}
	FailTaskStub        func(logger lager.Logger, taskGuid string, failureReason string) error
	failTaskMutex       sync.RWMutex
	failTaskArgsForCall []struct {
		logger        lager.Logger
		taskGuid      string
		failureReason string
	}
	failTaskReturns struct {
		result1 error
	}
	NewAuctioneerLockStub        func(auctioneerPresence models.AuctioneerPresence, interval time.Duration) (ifrit.Runner, error)
	newAuctioneerLockMutex       sync.RWMutex
	newAuctioneerLockArgsForCall []struct {
		auctioneerPresence models.AuctioneerPresence
		interval           time.Duration
	}
	newAuctioneerLockReturns struct {
		result1 ifrit.Runner
		result2 error
	}
	FailActualLRPStub        func(lager.Logger, models.ActualLRPKey, string) error
	failActualLRPMutex       sync.RWMutex
	failActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 string
	}
	failActualLRPReturns struct {
		result1 error
	}
}

func (fake *FakeAuctioneerBBS) Cells() ([]models.CellPresence, error) {
	fake.cellsMutex.Lock()
	fake.cellsArgsForCall = append(fake.cellsArgsForCall, struct{}{})
	fake.cellsMutex.Unlock()
	if fake.CellsStub != nil {
		return fake.CellsStub()
	} else {
		return fake.cellsReturns.result1, fake.cellsReturns.result2
	}
}

func (fake *FakeAuctioneerBBS) CellsCallCount() int {
	fake.cellsMutex.RLock()
	defer fake.cellsMutex.RUnlock()
	return len(fake.cellsArgsForCall)
}

func (fake *FakeAuctioneerBBS) CellsReturns(result1 []models.CellPresence, result2 error) {
	fake.CellsStub = nil
	fake.cellsReturns = struct {
		result1 []models.CellPresence
		result2 error
	}{result1, result2}
}

func (fake *FakeAuctioneerBBS) FailTask(logger lager.Logger, taskGuid string, failureReason string) error {
	fake.failTaskMutex.Lock()
	fake.failTaskArgsForCall = append(fake.failTaskArgsForCall, struct {
		logger        lager.Logger
		taskGuid      string
		failureReason string
	}{logger, taskGuid, failureReason})
	fake.failTaskMutex.Unlock()
	if fake.FailTaskStub != nil {
		return fake.FailTaskStub(logger, taskGuid, failureReason)
	} else {
		return fake.failTaskReturns.result1
	}
}

func (fake *FakeAuctioneerBBS) FailTaskCallCount() int {
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	return len(fake.failTaskArgsForCall)
}

func (fake *FakeAuctioneerBBS) FailTaskArgsForCall(i int) (lager.Logger, string, string) {
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	return fake.failTaskArgsForCall[i].logger, fake.failTaskArgsForCall[i].taskGuid, fake.failTaskArgsForCall[i].failureReason
}

func (fake *FakeAuctioneerBBS) FailTaskReturns(result1 error) {
	fake.FailTaskStub = nil
	fake.failTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLock(auctioneerPresence models.AuctioneerPresence, interval time.Duration) (ifrit.Runner, error) {
	fake.newAuctioneerLockMutex.Lock()
	fake.newAuctioneerLockArgsForCall = append(fake.newAuctioneerLockArgsForCall, struct {
		auctioneerPresence models.AuctioneerPresence
		interval           time.Duration
	}{auctioneerPresence, interval})
	fake.newAuctioneerLockMutex.Unlock()
	if fake.NewAuctioneerLockStub != nil {
		return fake.NewAuctioneerLockStub(auctioneerPresence, interval)
	} else {
		return fake.newAuctioneerLockReturns.result1, fake.newAuctioneerLockReturns.result2
	}
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLockCallCount() int {
	fake.newAuctioneerLockMutex.RLock()
	defer fake.newAuctioneerLockMutex.RUnlock()
	return len(fake.newAuctioneerLockArgsForCall)
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLockArgsForCall(i int) (models.AuctioneerPresence, time.Duration) {
	fake.newAuctioneerLockMutex.RLock()
	defer fake.newAuctioneerLockMutex.RUnlock()
	return fake.newAuctioneerLockArgsForCall[i].auctioneerPresence, fake.newAuctioneerLockArgsForCall[i].interval
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLockReturns(result1 ifrit.Runner, result2 error) {
	fake.NewAuctioneerLockStub = nil
	fake.newAuctioneerLockReturns = struct {
		result1 ifrit.Runner
		result2 error
	}{result1, result2}
}

func (fake *FakeAuctioneerBBS) FailActualLRP(arg1 lager.Logger, arg2 models.ActualLRPKey, arg3 string) error {
	fake.failActualLRPMutex.Lock()
	fake.failActualLRPArgsForCall = append(fake.failActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 models.ActualLRPKey
		arg3 string
	}{arg1, arg2, arg3})
	fake.failActualLRPMutex.Unlock()
	if fake.FailActualLRPStub != nil {
		return fake.FailActualLRPStub(arg1, arg2, arg3)
	} else {
		return fake.failActualLRPReturns.result1
	}
}

func (fake *FakeAuctioneerBBS) FailActualLRPCallCount() int {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return len(fake.failActualLRPArgsForCall)
}

func (fake *FakeAuctioneerBBS) FailActualLRPArgsForCall(i int) (lager.Logger, models.ActualLRPKey, string) {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return fake.failActualLRPArgsForCall[i].arg1, fake.failActualLRPArgsForCall[i].arg2, fake.failActualLRPArgsForCall[i].arg3
}

func (fake *FakeAuctioneerBBS) FailActualLRPReturns(result1 error) {
	fake.FailActualLRPStub = nil
	fake.failActualLRPReturns = struct {
		result1 error
	}{result1}
}

var _ bbs.AuctioneerBBS = new(FakeAuctioneerBBS)
