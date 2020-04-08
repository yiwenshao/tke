package logagent

import (
	"errors"
	"fmt"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	"sync"
	"time"
	clientset "tkestack.io/tke/api/client/clientset/versioned"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	logagentv1informer "tkestack.io/tke/api/client/informers/externalversions/logagent/v1"
	logagentv1lister "tkestack.io/tke/api/client/listers/logagent/v1"
	v1 "tkestack.io/tke/api/logagent/v1"
	controllerutil "tkestack.io/tke/pkg/controller"
	"tkestack.io/tke/pkg/logagent/util"

	//"tkestack.io/tke/pkg/platform/util"
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
	platformClient platformversionedclient.PlatformV1Interface
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
func NewController(platformClient platformversionedclient.PlatformV1Interface,client clientset.Interface,
	informer logagentv1informer.LogAgentInformer,
	resyncPeriod time.Duration) *Controller {
	// create the controller so we can inject the enqueue function
	controller := &Controller{
		platformClient:  platformClient,
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
	key, err := controllerutil.KeyFunc(obj)
	if err != nil {
		log.Error("Couldn't get key for LogCollector object",
			log.Any("object", obj), log.Err(err))
		return
	}
	c.queue.Add(key)
	log.Infof("enqueue logagent with key %v", key)
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

	err := c.syncLogCollector(key.(string))
	log.Infof("get key with name %v queue size is %v", key, c.queue.Len())
	if err == nil {
		c.queue.Forget(key)
		return true
	}

	runtime.HandleError(fmt.Errorf("error processing Logagent %s (will retry): %v", key, err))
	c.queue.AddRateLimited(key)


	return true
}


// syncLogCollector will sync the LogCollector with the given key if it has had
// its expectations fulfilled, meaning it did not expect to see any more of its
// namespaces created or deleted. This function is not meant to be invoked
// concurrently with the same key.
func (c *Controller) syncLogCollector(key string) error {
	startTime := time.Now()
	defer func() {
		log.Info("Finished syncing LogCollector", log.String("name", key), log.Duration("processTime", time.Since(startTime)))
	}()

	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	// LogCollector holds the latest LogCollector info from apiserver.
	LogCollector, err := c.lister.Get(name)
	switch {
	case k8serrors.IsNotFound(err):
		log.Info("LogCollector has been deleted. Attempting to cleanup resources", log.String("name", key))
		err = c.processLogCollectorDeletion(key)
	case err != nil:
		log.Warn("Unable to retrieve LogCollector from store", log.String("name", key), log.Err(err))
	default:
		cachedLogCollector := c.cache.getOrCreate(key)
		err = c.processLogCollectorUpdate(cachedLogCollector, LogCollector, key)
	}

	return err
}

func (c *Controller) processLogCollectorDeletion(key string) error {
	cachedLogCollector, ok := c.cache.get(key)
	if !ok {
		log.Error("LogCollector not in cache even though the watcher thought it was. Ignoring the deletion", log.String("name", key))
		return nil
	}
	return c.processLogCollectorDelete(cachedLogCollector, key)
}

func (c *Controller) processLogCollectorDelete(cachedLogCollector *cachedLogCollector, key string) error {
	log.Info("LogCollector will be dropped", log.String("name", key))

	if c.cache.Exist(key) {
		log.Info("Delete the LogCollector cache", log.String("name", key))
		c.cache.delete(key)
	}

	if _, ok := c.health.Load(key); ok {
		log.Info("Delete the LogCollector health cache", log.String("name", key))
		c.health.Delete(key)
	}

	LogCollector := cachedLogCollector.state
	return c.uninstallLogCollector(LogCollector)
}

func (c *Controller) uninstallLogCollector(LogCollector *v1.LogAgent) error {
	log.Info("Start to uninstall LogCollector",
		log.String("name", LogCollector.Name),
		log.String("clusterName", LogCollector.Spec.ClusterName))

	kubeClient, err := util.GetClusterClient(LogCollector.Spec.ClusterName,c.platformClient)
	if err != nil {
		log.Errorf("unable to get cluster client")
		return err
	}
	//cluster, err := c.client.PlatformV1().Clusters().Get(LogCollector.Spec.ClusterName, metav1.GetOptions{})
	//if err != nil && k8serrors.IsNotFound(err) {
	//	return nil
	//}
	//if err != nil {
	//	return err
	//}
	//kubeClient, err := util.BuildExternalClientSet(cluster, c.client.PlatformV1())
	//if err != nil {
	//	return err
	//}

	// Delete the operator daemonSet.
	clearDaemonSetErr := kubeClient.AppsV1().
		DaemonSets(metav1.NamespaceSystem).Delete(daemonSetName, &metav1.DeleteOptions{})
	// Delete the ClusterRoleBinding.
	clearCRBErr := kubeClient.RbacV1().
		ClusterRoleBindings().Delete(crbName, &metav1.DeleteOptions{})
	// Delete the ServiceAccount.
	clearSVCErr := kubeClient.CoreV1().ServiceAccounts(metav1.NamespaceSystem).
		Delete(svcAccountName, &metav1.DeleteOptions{})

	failed := false

	if clearDaemonSetErr != nil && !k8serrors.IsNotFound(clearDaemonSetErr) {
		failed = true
		log.Error("delete daemonSet for LogCollector failed",
			log.String("name", LogCollector.Name),
			log.String("clusterName", LogCollector.Spec.ClusterName),
			log.Err(clearDaemonSetErr))
	}

	if clearCRBErr != nil && !k8serrors.IsNotFound(clearCRBErr) {
		failed = true
		log.Error("delete crb for LogCollector failed",
			log.String("name", LogCollector.Name),
			log.String("clusterName", LogCollector.Spec.ClusterName),
			log.Err(clearCRBErr))
	}

	if clearSVCErr != nil && !k8serrors.IsNotFound(clearSVCErr) {
		failed = true
		log.Error("delete service account for LogCollector failed",
			log.String("name", LogCollector.Name),
			log.String("clusterName", LogCollector.Spec.ClusterName),
			log.Err(clearSVCErr))
	}

	if failed {
		return errors.New("delete LogCollector failed")
	}

	return nil
}



func (c *Controller) processLogCollectorUpdate(cachedLogCollector *cachedLogCollector, LogCollector *v1.LogAgent, key string) error {
	if cachedLogCollector.state != nil {
		// exist and the cluster name changed
		if cachedLogCollector.state.UID != LogCollector.UID {
			if err := c.processLogCollectorDelete(cachedLogCollector, key); err != nil {
				return err
			}
		}
	}
	err := c.createLogCollectorIfNeeded(key, cachedLogCollector, LogCollector)
	if err != nil {
		return err
	}

	cachedLogCollector.state = LogCollector
	// Always update the cache upon success.
	c.cache.set(key, cachedLogCollector)
	return nil
}


func (c *Controller) createLogCollectorIfNeeded(
	key string,
	cachedLogCollector *cachedLogCollector,
	LogCollector *v1.LogAgent) error {
	log.Info("Start to uninstall LogCollector",
		log.String("name", LogCollector.Name),
		log.String("clusterName", LogCollector.Spec.ClusterName))

	kubeClient, err := util.GetClusterClient(LogCollector.Spec.ClusterName,c.platformClient)
	if err != nil {
		log.Errorf("unable to get cluster client")
		return err
	}
	pod, err := kubeClient.CoreV1().Pods("default").Get("website", metav1.GetOptions{})
	if err != nil {
		log.Errorf("unable to get pods %v", err)
		return nil
	}
	log.Infof("get pod with name %v", pod.Name)
	return  nil
}
