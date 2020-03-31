package storage

import (
	"context"
	"fmt"
	"k8s.io/klog"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"tkestack.io/tke/api/logagent"
	"tkestack.io/tke/pkg/apiserver/authentication"
)

// TokenREST implements the REST endpoint.
type TokenREST struct {
	apiKeyStore *registry.Store

}

var _ = rest.Creater(&TokenREST{})

// New returns an empty object that can be used with Create after request data
// has been put into it.
func (r *TokenREST) New() runtime.Object {
	return &logagent.APIKeyReq{}
}


func (r *TokenREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	userName, tenantID := authentication.GetUsernameAndTenantID(ctx)

	apikeyReq := obj.(*logagent.APIKeyReq)

	klog.Infof("userName=%, tenantID=%v, apikeyReq=%+v spec=%+v", userName, tenantID, apikeyReq.Description, apikeyReq.Spec)
	var err error
	err = fmt.Errorf("test token, not implemented")

	return nil, apierrors.NewBadRequest(err.Error())
}

