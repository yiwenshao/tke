package storage

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	"tkestack.io/tke/api/logagent"
	"tkestack.io/tke/pkg/apiserver/authentication"
	"tkestack.io/tke/pkg/logagent/util"
	"tkestack.io/tke/pkg/util/log"
)
// TokenREST implements the REST endpoint.
type FileContentREST struct {
	//apiKeyStore *registry.Store
	//rest.Storage
	apiKeyStore *registry.Store
	PlatformClient platformversionedclient.PlatformV1Interface
	//*registry.Store
}

var _ = rest.Creater(&FileContentREST{})

func (r *FileContentREST)  New() runtime.Object {
	return &logagent.LogFileContent{}
}

func (r *FileContentREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	//how to get the parent resource??
	log.Infof("create filenode called")
	userName, tenantID := authentication.GetUsernameAndTenantID(ctx)
	fileContent := obj.(*logagent.LogFileContent)
	log.Infof("get userNmae %v tenantId %v and fileNode spec=%+v", userName, tenantID, fileContent.Spec)
	hostIp, err := util.GetClusterPodIp(fileContent.Spec.ClusterId, fileContent.Spec.Namespace, fileContent.Spec.Name, r.PlatformClient)
	if err != nil {
		return nil, errors.NewInternalError(fmt.Errorf("unable to get pod ip %v", err))
	}
	return &util.LocationStreamer{
		Request: util.FileNodeRequest{fileContent.Spec.Pod, fileContent.Spec.Namespace, fileContent.Spec.Container},
		Transport: nil,
		ContentType:     "text/plain",
		Ip: hostIp,
	}, nil
}



