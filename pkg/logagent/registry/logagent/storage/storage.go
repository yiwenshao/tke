package storage

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"tkestack.io/tke/api/logagent"
	"tkestack.io/tke/pkg/apiserver/authentication"
	apiserverutil "tkestack.io/tke/pkg/apiserver/util"
	registrylogagent "tkestack.io/tke/pkg/logagent/registry/logagent"
	"tkestack.io/tke/pkg/logagent/util"
	"tkestack.io/tke/pkg/util/log"
)

type Storage struct {
	LogAgent *REST
	LogFileTree *FileNodeREST
	Token     *TokenREST
	Status   *StatusREST
}

// NewStorage returns a Storage object that will work against channels.
func NewStorage(optsGetter genericregistry.RESTOptionsGetter, privilegedUsername string) *Storage {
	strategy := registrylogagent.NewStrategy()

	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &logagent.LogAgent{} },
		NewListFunc:              func() runtime.Object { return &logagent.LogAgentList{} },
		DefaultQualifiedResource: logagent.Resource("logagents"),
		PredicateFunc:            registrylogagent.MatchLogCollector,

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
		ExportStrategy: strategy,
	}

	options := &genericregistry.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc:    registrylogagent.GetAttrs,
	}

	if err := store.CompleteWithOptions(options); err != nil {
		log.Panic("Failed to create logagent etcd rest storage", log.Err(err))
	}

	statusStore := *store
	statusStore.UpdateStrategy = registrylogagent.NewStatusStrategy(strategy)
	statusStore.ExportStrategy = registrylogagent.NewStatusStrategy(strategy)

	return &Storage{
		LogAgent:     &REST{store, privilegedUsername},
		LogFileTree:  &FileNodeREST{store},
		Token: &TokenREST{store},
		Status:       &StatusREST{&statusStore},
	}

}

// ValidateGetObjectAndTenantID validate name and tenantID, if success return LogCollector
func ValidateGetObjectAndTenantID(ctx context.Context, store *registry.Store, name string, options *metav1.GetOptions) (runtime.Object, error) {
	obj, err := store.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}

	logCollector := obj.(*logagent.LogAgent)
	if err := util.FilterLogAgent(ctx, logCollector); err != nil {
		return nil, err
	}
	return logCollector, nil
}

// ValidateExportObjectAndTenantID validate name and tenantID, if success return LogCollector
func ValidateExportObjectAndTenantID(ctx context.Context, store *registry.Store, name string, options metav1.ExportOptions) (runtime.Object, error) {
	obj, err := store.Export(ctx, name, options)
	if err != nil {
		return nil, err
	}
	logCollector := obj.(*logagent.LogAgent)
	if err := util.FilterLogAgent(ctx, logCollector); err != nil {
		return nil, err
	}
	return logCollector, nil
}


// REST implements a RESTStorage for channels against etcd.
type REST struct {
	*registry.Store
	privilegedUsername string
}


var _ rest.ShortNamesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of
// short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"loga"}
}

//No need to implement TODO: remove this function
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error){
	la := obj.(*logagent.LogAgent)
	log.Infof("create logagents name=%v",la.Name)
	return r.Store.Create(ctx, obj, createValidation, options)
}


// List selects resources in the storage which match to the selector. 'options' can be nil.
func (r *REST) List(ctx context.Context, options *metainternal.ListOptions) (runtime.Object, error) {
	log.Infof("list logagents ")
	wrappedOptions := apiserverutil.PredicateListOptions(ctx, options)
	return r.Store.List(ctx, wrappedOptions)
}

// Get finds a resource in the storage by name and returns it.
func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	log.Infof("get logagents with name %v", name)
	return ValidateGetObjectAndTenantID(ctx, r.Store, name, options)
}

// Export an object.  Fields that are not user specified are stripped out
// Returns the stripped object.
func (r *REST) Export(ctx context.Context, name string, options metav1.ExportOptions) (runtime.Object, error) {
	log.Infof("export logagents with name %v", name)
	return ValidateExportObjectAndTenantID(ctx, r.Store, name, options)
}

// Update finds a resource in the storage and updates it.
func (r *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	log.Infof("update logagents with name %v", name)
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	// sub resources should never allow create on update.
	_, err := ValidateGetObjectAndTenantID(ctx, r.Store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	return r.Store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

// Delete enforces life-cycle rules for cluster termination
func (r *REST) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	log.Infof("delete logagents with name %v", name)
	_, err := ValidateGetObjectAndTenantID(ctx, r.Store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	return r.Store.Delete(ctx, name, deleteValidation, options)
}

// DeleteCollection selects all resources in the storage matching given 'listOptions'
// and deletes them.
func (r *REST) DeleteCollection(ctx context.Context, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions, listOptions *metainternal.ListOptions) (runtime.Object, error) {
	if !authentication.IsAdministrator(ctx, r.privilegedUsername) {
		return nil, errors.NewMethodNotSupported(logagent.Resource("logagents"), "delete collection")
	}
	return r.Store.DeleteCollection(ctx, deleteValidation, options, listOptions)
}



// StatusREST implements the REST endpoint for changing the status of a LogAgent.
type StatusREST struct {
	store *registry.Store
}

// StatusREST implements Patcher.
var _ = rest.Patcher(&StatusREST{})

// New returns an empty object that can be used with Create and Update after
// request data has been put into it.
func (r *StatusREST) New() runtime.Object {
	return r.store.New()
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	log.Infof("get logagents status name=%v", name)
	return ValidateGetObjectAndTenantID(ctx, r.store, name, options)
}

// Export an object.  Fields that are not user specified are stripped out
// Returns the stripped object.
func (r *StatusREST) Export(ctx context.Context, name string, options metav1.ExportOptions) (runtime.Object, error) {
	log.Infof("export logagents status name=%v", name)
	return ValidateExportObjectAndTenantID(ctx, r.store, name, options)
}

// Update alters the status subset of an object.
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	log.Infof("update logagents status name=%v", name)
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	// sub resources should never allow create on update.
	_, err := ValidateGetObjectAndTenantID(ctx, r.store, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

