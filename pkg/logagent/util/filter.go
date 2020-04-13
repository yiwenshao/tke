package util

import (
	"context"
	"tkestack.io/tke/api/logagent"
	"tkestack.io/tke/api/logagent/v1"
	"tkestack.io/tke/pkg/apiserver/authentication"
	"k8s.io/apimachinery/pkg/api/errors"
)

// FilterLogAgent is used to filter log collector that do not belong
// to the tenant.
func FilterLogAgent(ctx context.Context, decorator *logagent.LogAgent) error {
	_, tenantID := authentication.GetUsernameAndTenantID(ctx)
	if tenantID == "" {
		return nil
	}
	if decorator.Spec.TenantID != tenantID {
		return errors.NewNotFound(v1.Resource("logagent"), decorator.ObjectMeta.Name)
	}
	return nil
}
