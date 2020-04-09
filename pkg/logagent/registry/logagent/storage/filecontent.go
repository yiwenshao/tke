package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"net/http"
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


type FileContentProxy struct {
	Req logagent.LogFileContentSpec
	Ip string
	Port string
}

func (p *FileContentProxy) GetReaderCloser() io.ReadCloser {
	jsonStr, err := json.Marshal(p.Req)
	if err != nil {
		log.Errorf("unable to marshal request to json %v", err)
		return nil
	}
	url := "http://" + p.Ip + ":" + p.Port + "/v1/logfile/content"
	log.Infof("url is %v", url)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	httpReq.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Errorf("unable to generate request %v", err)
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Errorf("unable to connect to log-agent %v", err)
		return nil
	}
	return resp.Body
}


func (r *FileContentREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	//how to get the parent resource??
	log.Infof("create filenode called")
	userName, tenantID := authentication.GetUsernameAndTenantID(ctx)
	fileContent := obj.(*logagent.LogFileContent)
	log.Infof("get userNmae %v tenantId %v and fileNode spec=%+v", userName, tenantID, fileContent.Spec)
	hostIp, err := util.GetClusterPodIp(fileContent.Spec.ClusterId, fileContent.Spec.Namespace, fileContent.Spec.Pod, r.PlatformClient)
	if err != nil {
		return nil, errors.NewInternalError(fmt.Errorf("unable to get pod ip %v", err))
	}
	return &util.LocationStreamer{
		Request: &FileContentProxy{Req:fileContent.Spec,Ip:hostIp,Port:util.LogagentPort},
		Transport: nil,
		ContentType: "application/json",
		Ip: hostIp,
	}, nil
}



