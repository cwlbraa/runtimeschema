// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
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
	NewAuctioneerLockStub        func(auctioneerPresence models.AuctioneerPresence, retryInterval time.Duration) (ifrit.Runner, error)
	newAuctioneerLockMutex       sync.RWMutex
	newAuctioneerLockArgsForCall []struct {
		auctioneerPresence models.AuctioneerPresence
		retryInterval      time.Duration
	}
	newAuctioneerLockReturns struct {
		result1 ifrit.Runner
		result2 error
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

func (fake *FakeAuctioneerBBS) NewAuctioneerLock(auctioneerPresence models.AuctioneerPresence, retryInterval time.Duration) (ifrit.Runner, error) {
	fake.newAuctioneerLockMutex.Lock()
	fake.newAuctioneerLockArgsForCall = append(fake.newAuctioneerLockArgsForCall, struct {
		auctioneerPresence models.AuctioneerPresence
		retryInterval      time.Duration
	}{auctioneerPresence, retryInterval})
	fake.newAuctioneerLockMutex.Unlock()
	if fake.NewAuctioneerLockStub != nil {
		return fake.NewAuctioneerLockStub(auctioneerPresence, retryInterval)
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
	return fake.newAuctioneerLockArgsForCall[i].auctioneerPresence, fake.newAuctioneerLockArgsForCall[i].retryInterval
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLockReturns(result1 ifrit.Runner, result2 error) {
	fake.NewAuctioneerLockStub = nil
	fake.newAuctioneerLockReturns = struct {
		result1 ifrit.Runner
		result2 error
	}{result1, result2}
}

var _ bbs.AuctioneerBBS = new(FakeAuctioneerBBS)
