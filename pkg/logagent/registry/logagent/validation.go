package logagent

import (
	apiMachineryValidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"tkestack.io/tke/api/logagent"
)

// ValidateName is a ValidateNameFunc for names that must be a DNS sub domain.
var ValidateName = apiMachineryValidation.ValidateNamespaceName

// ValidateLogCollector tests if required fields in the cluster are set.
func ValidateLogCollector(decorator *logagent.LogAgent) field.ErrorList {
	allErrs := apiMachineryValidation.ValidateObjectMeta(&decorator.ObjectMeta, false, ValidateName, field.NewPath("metadata"))

	if len(decorator.Spec.ClusterName) == 0 {
		allErrs = append(allErrs, field.Required(field.NewPath("spec", "clusterName"), "must specify a cluster name"))
	}

	return allErrs
}

// ValidateLogCollectorUpdate tests if required fields in the namespace set are
// set during an update.
func ValidateLogCollectorUpdate(new *logagent.LogAgent, old *logagent.LogAgent) field.ErrorList {
	allErrs := apiMachineryValidation.ValidateObjectMetaUpdate(&new.ObjectMeta, &old.ObjectMeta, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidateLogCollector(new)...)

	if new.Spec.ClusterName != old.Spec.ClusterName {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "clusterName"), new.Spec.ClusterName, "disallowed change the cluster name"))
	}

	if new.Spec.TenantID != old.Spec.TenantID {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "tenantID"), new.Spec.TenantID, "disallowed change the tenant"))
	}

	if new.Status.Phase == "" {
		allErrs = append(allErrs, field.Required(field.NewPath("status", "phase"), string(new.Status.Phase)))
	}

	return allErrs
}
