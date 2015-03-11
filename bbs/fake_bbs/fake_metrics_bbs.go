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

type FakeMetricsBBS struct {
	TasksStub        func(logger lager.Logger) ([]models.Task, error)
	tasksMutex       sync.RWMutex
	tasksArgsForCall []struct {
		logger lager.Logger
	}
	tasksReturns struct {
		result1 []models.Task
		result2 error
	}
	ServiceRegistrationsStub        func() (models.ServiceRegistrations, error)
	serviceRegistrationsMutex       sync.RWMutex
	serviceRegistrationsArgsForCall []struct{}
	serviceRegistrationsReturns struct {
		result1 models.ServiceRegistrations
		result2 error
	}
	DomainsStub        func() ([]string, error)
	domainsMutex       sync.RWMutex
	domainsArgsForCall []struct{}
	domainsReturns struct {
		result1 []string
		result2 error
	}
	DesiredLRPsStub        func() ([]models.DesiredLRP, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct{}
	desiredLRPsReturns struct {
		result1 []models.DesiredLRP
		result2 error
	}
	ActualLRPsStub        func() ([]models.ActualLRP, error)
	actualLRPsMutex       sync.RWMutex
	actualLRPsArgsForCall []struct{}
	actualLRPsReturns struct {
		result1 []models.ActualLRP
		result2 error
	}
	NewRuntimeMetricsLockStub        func(runtimeMetricsID string, interval time.Duration) ifrit.Runner
	newRuntimeMetricsLockMutex       sync.RWMutex
	newRuntimeMetricsLockArgsForCall []struct {
		runtimeMetricsID string
		interval         time.Duration
	}
	newRuntimeMetricsLockReturns struct {
		result1 ifrit.Runner
	}
}

func (fake *FakeMetricsBBS) Tasks(logger lager.Logger) ([]models.Task, error) {
	fake.tasksMutex.Lock()
	fake.tasksArgsForCall = append(fake.tasksArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.tasksMutex.Unlock()
	if fake.TasksStub != nil {
		return fake.TasksStub(logger)
	} else {
		return fake.tasksReturns.result1, fake.tasksReturns.result2
	}
}

func (fake *FakeMetricsBBS) TasksCallCount() int {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return len(fake.tasksArgsForCall)
}

func (fake *FakeMetricsBBS) TasksArgsForCall(i int) lager.Logger {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return fake.tasksArgsForCall[i].logger
}

func (fake *FakeMetricsBBS) TasksReturns(result1 []models.Task, result2 error) {
	fake.TasksStub = nil
	fake.tasksReturns = struct {
		result1 []models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeMetricsBBS) ServiceRegistrations() (models.ServiceRegistrations, error) {
	fake.serviceRegistrationsMutex.Lock()
	fake.serviceRegistrationsArgsForCall = append(fake.serviceRegistrationsArgsForCall, struct{}{})
	fake.serviceRegistrationsMutex.Unlock()
	if fake.ServiceRegistrationsStub != nil {
		return fake.ServiceRegistrationsStub()
	} else {
		return fake.serviceRegistrationsReturns.result1, fake.serviceRegistrationsReturns.result2
	}
}

func (fake *FakeMetricsBBS) ServiceRegistrationsCallCount() int {
	fake.serviceRegistrationsMutex.RLock()
	defer fake.serviceRegistrationsMutex.RUnlock()
	return len(fake.serviceRegistrationsArgsForCall)
}

func (fake *FakeMetricsBBS) ServiceRegistrationsReturns(result1 models.ServiceRegistrations, result2 error) {
	fake.ServiceRegistrationsStub = nil
	fake.serviceRegistrationsReturns = struct {
		result1 models.ServiceRegistrations
		result2 error
	}{result1, result2}
}

func (fake *FakeMetricsBBS) Domains() ([]string, error) {
	fake.domainsMutex.Lock()
	fake.domainsArgsForCall = append(fake.domainsArgsForCall, struct{}{})
	fake.domainsMutex.Unlock()
	if fake.DomainsStub != nil {
		return fake.DomainsStub()
	} else {
		return fake.domainsReturns.result1, fake.domainsReturns.result2
	}
}

func (fake *FakeMetricsBBS) DomainsCallCount() int {
	fake.domainsMutex.RLock()
	defer fake.domainsMutex.RUnlock()
	return len(fake.domainsArgsForCall)
}

func (fake *FakeMetricsBBS) DomainsReturns(result1 []string, result2 error) {
	fake.DomainsStub = nil
	fake.domainsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeMetricsBBS) DesiredLRPs() ([]models.DesiredLRP, error) {
	fake.desiredLRPsMutex.Lock()
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct{}{})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub()
	} else {
		return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
	}
}

func (fake *FakeMetricsBBS) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeMetricsBBS) DesiredLRPsReturns(result1 []models.DesiredLRP, result2 error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeMetricsBBS) ActualLRPs() ([]models.ActualLRP, error) {
	fake.actualLRPsMutex.Lock()
	fake.actualLRPsArgsForCall = append(fake.actualLRPsArgsForCall, struct{}{})
	fake.actualLRPsMutex.Unlock()
	if fake.ActualLRPsStub != nil {
		return fake.ActualLRPsStub()
	} else {
		return fake.actualLRPsReturns.result1, fake.actualLRPsReturns.result2
	}
}

func (fake *FakeMetricsBBS) ActualLRPsCallCount() int {
	fake.actualLRPsMutex.RLock()
	defer fake.actualLRPsMutex.RUnlock()
	return len(fake.actualLRPsArgsForCall)
}

func (fake *FakeMetricsBBS) ActualLRPsReturns(result1 []models.ActualLRP, result2 error) {
	fake.ActualLRPsStub = nil
	fake.actualLRPsReturns = struct {
		result1 []models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeMetricsBBS) NewRuntimeMetricsLock(runtimeMetricsID string, interval time.Duration) ifrit.Runner {
	fake.newRuntimeMetricsLockMutex.Lock()
	fake.newRuntimeMetricsLockArgsForCall = append(fake.newRuntimeMetricsLockArgsForCall, struct {
		runtimeMetricsID string
		interval         time.Duration
	}{runtimeMetricsID, interval})
	fake.newRuntimeMetricsLockMutex.Unlock()
	if fake.NewRuntimeMetricsLockStub != nil {
		return fake.NewRuntimeMetricsLockStub(runtimeMetricsID, interval)
	} else {
		return fake.newRuntimeMetricsLockReturns.result1
	}
}

func (fake *FakeMetricsBBS) NewRuntimeMetricsLockCallCount() int {
	fake.newRuntimeMetricsLockMutex.RLock()
	defer fake.newRuntimeMetricsLockMutex.RUnlock()
	return len(fake.newRuntimeMetricsLockArgsForCall)
}

func (fake *FakeMetricsBBS) NewRuntimeMetricsLockArgsForCall(i int) (string, time.Duration) {
	fake.newRuntimeMetricsLockMutex.RLock()
	defer fake.newRuntimeMetricsLockMutex.RUnlock()
	return fake.newRuntimeMetricsLockArgsForCall[i].runtimeMetricsID, fake.newRuntimeMetricsLockArgsForCall[i].interval
}

func (fake *FakeMetricsBBS) NewRuntimeMetricsLockReturns(result1 ifrit.Runner) {
	fake.NewRuntimeMetricsLockStub = nil
	fake.newRuntimeMetricsLockReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

var _ bbs.MetricsBBS = new(FakeMetricsBBS)
