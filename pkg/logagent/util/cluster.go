package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sync"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	"tkestack.io/tke/pkg/platform/util"
	"tkestack.io/tke/pkg/util/log"
)

// ClusterNameToClient mapping cluster to kubernetes client
// clusterName => kubernetes.Interface
var ClusterNameToClient sync.Map

// ClusterNameToMonitor mapping cluster to monitoring client
// clusterName => monitoringclient.Clientset
var ClusterNameToMonitor sync.Map

// GetClusterClient get kubernetes client via cluster name
func GetClusterClient(clusterName string, platformClient platformversionedclient.PlatformV1Interface) (kubernetes.Interface, error) {
	// First check from cache
	if item, ok := ClusterNameToClient.Load(clusterName); ok {
		// Check if is available
		kubeClient := item.(kubernetes.Interface)
		_, err := kubeClient.CoreV1().Services(metav1.NamespaceSystem).List(metav1.ListOptions{})
		if err == nil {
			return kubeClient, nil
		}
		ClusterNameToClient.Delete(clusterName)
	}

	kubeClient, err := util.BuildExternalClientSetWithName(platformClient, clusterName)
	if err != nil {
		return nil, err
	}

	ClusterNameToClient.Store(clusterName, kubeClient)

	return kubeClient, nil
}


//use cache to optimize this function
func GetClusterPodIp(clusterName, namespace, podName string, platformClient platformversionedclient.PlatformV1Interface) (string, error) {
	client, err := GetClusterClient(clusterName, platformClient)
	if err != nil {
		log.Errorf("unable to get cluster client %v", err)
		return "", err
	}
	pod, err := client.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("unable to get pod in cluster %v err=%v", clusterName, err)
		return "", err
	}
	return pod.Status.HostIP, nil
}
