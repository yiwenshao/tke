package logagent

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	"sync"
	"time"
	clientset "tkestack.io/tke/api/client/clientset/versioned"
	logagentv1informer "tkestack.io/tke/api/client/informers/externalversions/logagent/v1"
	logagentv1lister "tkestack.io/tke/api/client/listers/logagent/v1"
	v1 "tkestack.io/tke/api/logagent/v1"
	"tkestack.io/tke/pkg/util/log"
	"tkestack.io/tke/pkg/util/metrics"
)

const (
	controllerName = "logagent-controller"

	crbName        = "logagent-role-binding"
	svcAccountName = "logagent"
	daemonSetName  = "logagent"

	clientRetryCount    = 5
	clientRetryInterval = 5 * time.Second

	timeOut       = 5 * time.Minute
	maxRetryCount = 5

	upgradePatchTemplate = `[{"op":"replace","path":"/spec/template/spec/containers/0/image","value":"%s"}]`
)

// Controller is responsible for performing actions dependent upon a LogCollector phase.
type Controller struct {
	client       clientset.Interface
	cache        *logcollectorCache
	health       sync.Map
	checking     sync.Map
	upgrading    sync.Map
	queue        workqueue.RateLimitingInterface
	lister       logagentv1lister.LogAgentLister
	listerSynced cache.InformerSynced
	stopCh       <-chan struct{}
}


// NewController creates a new LogCollector Controller object.
func NewController(client clientset.Interface,
	informer logagentv1informer.LogAgentInformer,
	resyncPeriod time.Duration) *Controller {
	// create the controller so we can inject the enqueue function
	controller := &Controller{
		client: client,
		cache:  &logcollectorCache{lcMap: make(map[string]*cachedLogCollector)},
		queue:  workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerName),
	}

	if client != nil && client.LogagentV1().RESTClient().GetRateLimiter() != nil {
		_ = metrics.RegisterMetricAndTrackRateLimiterUsage(controllerName, client.LogagentV1().RESTClient().GetRateLimiter())
	}

	// configure the informer event handlers
	informer.Informer().AddEventHandlerWithResyncPeriod(
		cache.ResourceEventHandlerFuncs{
			AddFunc: controller.enqueueLogAgent,
			UpdateFunc: func(oldObj, newObj interface{}) {
				oldLogCollector, ok1 := oldObj.(*v1.LogAgent)
				curLogCollector, ok2 := newObj.(*v1.LogAgent)
				if ok1 && ok2 && controller.needsUpdate(oldLogCollector, curLogCollector) {
					controller.enqueueLogAgent(newObj)
				}
			},
			DeleteFunc: controller.enqueueLogAgent,
		},
		resyncPeriod,
	)
	controller.lister = informer.Lister()
	controller.listerSynced = informer.Informer().HasSynced

	return controller
}



func (c *Controller) enqueueLogAgent(obj interface{}) {
	log.Infof("enqueue logagent %+v", obj)
}


func (c *Controller) needsUpdate(old *v1.LogAgent, new *v1.LogAgent) bool {
	return !reflect.DeepEqual(old, new)
}


// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers.
func (c *Controller) Run(workers int, stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	log.Info("Starting Logagent controller")
	defer log.Info("Shutting down Logagent controller")

	if !cache.WaitForCacheSync(stopCh, c.listerSynced) {
		return fmt.Errorf("failed to wait for LogAgent cache to sync")
	}

	c.stopCh = stopCh

	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}

	<-stopCh
	return nil
}

// worker processes the queue of namespace objects.
// Each key can be in the queue at most once.
// The system ensures that no two workers can process
// the same namespace at the same time.
func (c *Controller) worker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	//err := c.syncLogCollector(key.(string))
	err := fmt.Errorf("")
	err = nil
	log.Infof("get key with name %v queue size is %v", key, c.queue.Len())
	if err == nil {
		c.queue.Forget(key)
		return true
	}

	runtime.HandleError(fmt.Errorf("error processing Logagent %s (will retry): %v", key, err))
	c.queue.AddRateLimited(key)
	return true
}

