package configmap

import (
	apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"tkestack.io/tke/api/logagent"
)

// ValidateName is a ValidateNameFunc for names that must be a DNS
// subdomain.
var ValidateName = apimachineryvalidation.ValidateNamespaceName

// ValidateConfigMap tests if required fields in the cluster are set.
func ValidateConfigMap(configmap *logagent.ConfigMap) field.ErrorList {
	allErrs := apimachineryvalidation.ValidateObjectMeta(&configmap.ObjectMeta, false, ValidateName, field.NewPath("metadata"))

	return allErrs
}

// ValidateConfigMapUpdate tests if required fields in the namespace set are
// set during an update.
func ValidateConfigMapUpdate(configmap *logagent.ConfigMap, old *logagent.ConfigMap) field.ErrorList {
	allErrs := apimachineryvalidation.ValidateObjectMetaUpdate(&configmap.ObjectMeta, &old.ObjectMeta, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidateConfigMap(configmap)...)

	return allErrs
}
