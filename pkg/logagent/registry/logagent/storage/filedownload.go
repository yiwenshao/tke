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
type FileDownloadREST struct {
	//apiKeyStore *registry.Store
	//rest.Storage
	apiKeyStore *registry.Store
	PlatformClient platformversionedclient.PlatformV1Interface
	//*registry.Store
}


var _ = rest.Creater(&FileDownloadREST{})

func (r *FileDownloadREST)  New() runtime.Object {
	return &logagent.LogFileDownload{}
}

func (r *FileDownloadREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	log.Infof("create filenode called")
	userName, tenantID := authentication.GetUsernameAndTenantID(ctx)
	fileDownload := obj.(*logagent.LogFileDownload)
	log.Infof("get userNmae %v tenantId %v and fileNode spec=%+v", userName, tenantID, fileDownload.Spec)
	hostIp, err := util.GetClusterPodIp(fileDownload.Spec.ClusterId, fileDownload.Spec.Namespace, fileDownload.Spec.Name, r.PlatformClient)
	if err != nil {
		return nil, errors.NewInternalError(fmt.Errorf("unable to get pod ip %v", err))
	}
	return &util.LocationStreamer{
		Request: util.FileNodeRequest{fileDownload.Spec.Pod, fileDownload.Spec.Namespace, fileDownload.Spec.Container},
		Transport: nil,
		ContentType:     "text/plain",
		Ip: hostIp,
	}, nil
}
