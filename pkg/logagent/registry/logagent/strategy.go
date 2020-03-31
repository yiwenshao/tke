package logagent

import (
	"context"
	"tkestack.io/tke/api/logagent"

	"tkestack.io/tke/pkg/apiserver/authentication"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"tkestack.io/tke/pkg/util/log"
	namesutil "tkestack.io/tke/pkg/util/names"
)

// Strategy implements verification logic for tapp controller.
type Strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var _ rest.RESTCreateStrategy = &Strategy{}
var _ rest.RESTUpdateStrategy = &Strategy{}
var _ rest.RESTDeleteStrategy = &Strategy{}

// NewStrategy creates a strategy that is the default logic that applies when
// creating and updating namespace set objects.
func NewStrategy() *Strategy {
	return &Strategy{logagent.Scheme, namesutil.Generator}
}

// DefaultGarbageCollectionPolicy returns the default garbage collection behavior.
func (Strategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	return rest.Unsupported
}

// NamespaceScoped is false for namespaceSets
func (Strategy) NamespaceScoped() bool {
	return false
}

// Export strips fields that can not be set by the user.
func (Strategy) Export(ctx context.Context, obj runtime.Object, exact bool) error {
	return nil
}

// PrepareForCreate is invoked on create before validation to normalize
// the object.
func (Strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_, tenantID := authentication.GetUsernameAndTenantID(ctx)
	logCollector, _ := obj.(*logagent.LogAgent)

	if len(tenantID) != 0 {
		logCollector.Spec.TenantID = tenantID
	}

	if logCollector.Name == "" && logCollector.GenerateName == "" {
		logCollector.GenerateName = "lc-"
	}

	if logCollector.Spec.Version == "" {
		logCollector.Spec.Version = "1.0.0" //TODO: replace with const
	}
}

// PrepareForUpdate is invoked on update before validation to normalize the
// object.
func (Strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_, tenantID := authentication.GetUsernameAndTenantID(ctx)
	oldCollector := old.(*logagent.LogAgent)
	newCollector, _ := obj.(*logagent.LogAgent)
	if len(tenantID) != 0 {
		if oldCollector.Spec.TenantID != tenantID {
			log.Panic("Unauthorized update new LogCollector information", log.String("oldTenantID", oldCollector.Spec.TenantID), log.String("newTenantID", newCollector.Spec.TenantID), log.String("userTenantID", tenantID))
		}
		newCollector.Spec.TenantID = tenantID
	}
	newCollector.Status = oldCollector.Status
}

// Validate validates a new tapp controller.
func (Strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return ValidateLogCollector(obj.(*logagent.LogAgent))
}

// AllowCreateOnUpdate is false for persistent events
func (Strategy) AllowCreateOnUpdate() bool {
	return false
}

// AllowUnconditionalUpdate returns true if the object can be updated
// unconditionally (irrespective of the latest resource version), when there is
// no resource version specified in the object.
func (Strategy) AllowUnconditionalUpdate() bool {
	return false
}

// Canonicalize normalizes the object after validation.
func (Strategy) Canonicalize(obj runtime.Object) {
}

// ValidateUpdate is the default update validation for an end namespace set.
func (Strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return ValidateLogCollectorUpdate(obj.(*logagent.LogAgent), old.(*logagent.LogAgent))
}

// GetAttrs returns labels and fields of a given object for filtering purposes.
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	logCollector, _ := obj.(*logagent.LogAgent)
	return labels.Set(logCollector.ObjectMeta.Labels), ToSelectableFields(logCollector), nil
}

// MatchLogCollector returns a generic matcher for a given label and field selector.
func MatchLogCollector(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
		IndexFields: []string{
			"spec.tenantID",
			"spec.clusterName",
			"spec.version",
			"status.version",
			"status.phase",
		},
	}
}

// ToSelectableFields returns a field set that represents the object
func ToSelectableFields(csiOperator *logagent.LogAgent) fields.Set {
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&csiOperator.ObjectMeta, false)
	specificFieldsSet := fields.Set{
		"spec.tenantID":    csiOperator.Spec.TenantID,
		"spec.clusterName": csiOperator.Spec.ClusterName,
		"spec.version":     csiOperator.Spec.Version,
		"status.version":   csiOperator.Status.Version,
		"status.phase":     string(csiOperator.Status.Phase),
	}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}

// StatusStrategy implements verification logic for status of tapp controller.
type StatusStrategy struct {
	*Strategy
}

var _ rest.RESTUpdateStrategy = &StatusStrategy{}

// NewStatusStrategy create the StatusStrategy object by given strategy.
func NewStatusStrategy(strategy *Strategy) *StatusStrategy {
	return &StatusStrategy{strategy}
}

// PrepareForUpdate is invoked on update before validation to normalize
// the object.  For example: remove fields that are not to be persisted,
// sort order-insensitive list fields, etc.  This should not remove fields
// whose presence would be considered a validation error.
func (StatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	newCollector := obj.(*logagent.LogAgent)
	oldCollector := old.(*logagent.LogAgent)
	newCollector.Spec = oldCollector.Spec
}

// ValidateUpdate is invoked after default fields in the object have been
// filled in before the object is persisted.  This method should not mutate
// the object.
func (StatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}