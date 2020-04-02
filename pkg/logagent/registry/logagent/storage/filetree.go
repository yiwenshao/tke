package storage

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	"tkestack.io/tke/api/logagent"
	"tkestack.io/tke/pkg/apiserver/authentication"
	"tkestack.io/tke/pkg/logagent/util"
	"tkestack.io/tke/pkg/util/log"
)

// TokenREST implements the REST endpoint.
type FileNodeREST struct {
	//apiKeyStore *registry.Store
	//rest.Storage
	apiKeyStore *registry.Store
	PlatformClient platformversionedclient.PlatformV1Interface
	//*registry.Store
}

var _ = rest.Creater(&FileNodeREST{})//implement the Creater interface, then how to obtail client to user cluster?
//var _ rest.ShortNamesProvider = &FileNodeREST{}
//var _ rest.Creater = &FileNodeREST{}
//var _ rest.Scoper = &FileNodeREST{}
////more interfaces to be created??

// New returns an empty object that can be used with Create after request data
// has been put into it.
func (r *FileNodeREST)  New() runtime.Object {
	log.Infof("new filenode called")
	return &logagent.LogFileTree{}
}

//func (r *FileNodeREST) ShortNames() []string{
//	return []string{"logfiletree"}
//}
//
//func (r *FileNodeREST) NamespaceScoped() bool {
//	return false
//}

type  String string

func (*String) GetObjectKind() schema.ObjectKind {
	return schema.EmptyObjectKind
}

func (*String) DeepCopyObject() runtime.Object {
	panic("String does not implement DeepCopyObject")
}

func (r *FileNodeREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	//how to get the parent resource??
	log.Infof("create filenode called")
	userName, tenantID := authentication.GetUsernameAndTenantID(ctx)
	fileNode := obj.(*logagent.LogFileTree)
	log.Infof("get userNmae %v tenantId %v and fileNode spec=%+v", userName, tenantID, fileNode.Spec)

	client, err := util.GetClusterClient(fileNode.Spec.ClusterId,r.PlatformClient)
	if err != nil {
		log.Infof("unable to connect to user cluster %v", err)
		return nil, errors.NewInternalError(fmt.Errorf("test to not implemented log filenode"))
	}
	_, err = client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Infof("unable to get cluster nodes")
	}
	//log.Infof("my nodes are %+v", nodes)

	pod, err := client.CoreV1().Pods(fileNode.Spec.Namespace).Get(fileNode.Spec.Pod,metav1.GetOptions{})
	if err != nil {
		log.Errorf("unable to get pod %v", err)
		return nil, errors.NewInternalError(fmt.Errorf("test to not implemented log filenode"))
	}
	hostip := pod.Status.HostIP
	res := util.GetPodFileTree( util.FileNodeRequest{fileNode.Spec.Pod, fileNode.Spec.Namespace, fileNode.Spec.Container}, hostip)

	log.Infof("get file node results %v", res)

	//return nil, nil
	//return &logagent.LogFileTree{
	//	Spec: logagent.LogFileTreeSpec{Container:string(res)},
	//}, nil

	return &util.LocationStreamer{
		Request: util.FileNodeRequest{fileNode.Spec.Pod, fileNode.Spec.Namespace, fileNode.Spec.Container},
		Transport: nil,
		ContentType:     "text/plain",
		Ip: hostip,
	}, nil

}
