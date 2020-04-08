package storage

import (
	"context"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"tkestack.io/tke/api/logagent"
	apiserverutil "tkestack.io/tke/pkg/apiserver/util"
	"tkestack.io/tke/pkg/logagent/registry/configmap"
	"tkestack.io/tke/pkg/util/log"
)

// Storage includes storage for configmap and all sub resources.
type Storage struct {
	ConfigMap *REST
}

// NewStorage returns a Storage object that will work against configmap.
func NewStorage(optsGetter genericregistry.RESTOptionsGetter) *Storage {
	strategy := configmap.NewStrategy()
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &logagent.ConfigMap{} },
		NewListFunc:              func() runtime.Object { return &logagent.ConfigMapList{} },
		DefaultQualifiedResource: logagent.Resource("configmaps"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
		ExportStrategy: strategy,
	}

	options := &genericregistry.StoreOptions{
		RESTOptions: optsGetter,
	}

	if err := store.CompleteWithOptions(options); err != nil {
		log.Panic("Failed to create configmap etcd rest storage", log.Err(err))
	}

	return &Storage{
		ConfigMap: &REST{store},
	}
}

// REST implements a RESTStorage for configmap against etcd.
type REST struct {
	*registry.Store
}

var _ rest.ShortNamesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"cm"}
}

// List selects resources in the storage which match to the selector. 'options' can be nil.
func (r *REST) List(ctx context.Context, options *metainternal.ListOptions) (runtime.Object, error) {
	wrappedOptions := apiserverutil.PredicateListOptions(ctx, options)
	return r.Store.List(ctx, wrappedOptions)
}
