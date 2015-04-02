package task_bbs_test

import (
	"github.com/cloudfoundry-incubator/consuladapter"
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/services_bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/shared"
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/task_bbs"
	cbfakes "github.com/cloudfoundry-incubator/runtime-schema/cb/fakes"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/cloudfoundry/storeadapter"
	"github.com/cloudfoundry/storeadapter/storerunner/etcdstorerunner"
	"github.com/hashicorp/consul/consul/structs"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/clock/fakeclock"
	"github.com/pivotal-golang/lager/lagertest"

	"testing"
	"time"
)

const receptorURL = "http://some-receptor-url"

var etcdRunner *etcdstorerunner.ETCDClusterRunner
var etcdClient storeadapter.StoreAdapter
var consulRunner *consuladapter.ClusterRunner
var consulAdapter consuladapter.Adapter
var logger *lagertest.TestLogger
var servicesBBS *services_bbs.ServicesBBS
var fakeTaskClient *cbfakes.FakeTaskClient
var fakeAuctioneerClient *cbfakes.FakeAuctioneerClient
var fakeCellClient *cbfakes.FakeCellClient
var clock *fakeclock.FakeClock
var bbs *task_bbs.TaskBBS

var dummyAction = &models.RunAction{
	Path: "cat",
	Args: []string{"/tmp/file"},
}

func TestTaskBbs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Task BBS Suite")
}

var _ = BeforeSuite(func() {
	etcdRunner = etcdstorerunner.NewETCDClusterRunner(5001+config.GinkgoConfig.ParallelNode, 1)
	etcdClient = etcdRunner.RetryableAdapter()
	etcdRunner.Start()

	consulRunner = consuladapter.NewClusterRunner(
		9001+config.GinkgoConfig.ParallelNode*consuladapter.PortOffsetLength,
		1,
		"http",
	)
	consulAdapter = consulRunner.NewAdapter()
	consulRunner.Start()
})

var _ = AfterSuite(func() {
	etcdClient.Disconnect()
	etcdRunner.Stop()

	consulRunner.Stop()
})

var _ = BeforeEach(func() {
	etcdRunner.Reset()

	consulRunner.WaitUntilReady()
	consulRunner.Reset()

	logger = lagertest.NewTestLogger("test")

	fakeTaskClient = new(cbfakes.FakeTaskClient)
	fakeAuctioneerClient = new(cbfakes.FakeAuctioneerClient)
	fakeCellClient = new(cbfakes.FakeCellClient)
	clock = fakeclock.NewFakeClock(time.Unix(1238, 0))
	servicesBBS = services_bbs.New(consulAdapter, clock, logger)
	bbs = task_bbs.New(etcdClient, consulAdapter, clock, fakeTaskClient, fakeAuctioneerClient, fakeCellClient,
		servicesBBS, receptorURL)
})

func registerAuctioneer(auctioneer models.AuctioneerPresence) {
	jsonBytes, err := models.ToJSON(auctioneer)
	Ω(err).ShouldNot(HaveOccurred())

	cancelChan := make(chan struct{})
	_, err = consulAdapter.AcquireAndMaintainLock(
		shared.LockSchemaPath("auctioneer_lock"),
		jsonBytes,
		structs.SessionTTLMin,
		cancelChan)
	Ω(err).ShouldNot(HaveOccurred())
}

func registerCell(cell models.CellPresence) {
	jsonBytes, err := models.ToJSON(cell)
	Ω(err).ShouldNot(HaveOccurred())

	cancelChan := make(chan struct{})
	_, err = consulAdapter.AcquireAndMaintainLock(
		shared.CellSchemaPath(cell.CellID),
		jsonBytes,
		structs.SessionTTLMin,
		cancelChan)
	Ω(err).ShouldNot(HaveOccurred())
}
