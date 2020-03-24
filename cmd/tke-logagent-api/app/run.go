package app

import (
	"tkestack.io/tke/cmd/tke-logagent-api/app/config"
	"tkestack.io/tke/pkg/util/log"
)

// Run runs the specified TKE console. This should never exit.
func Run(cfg *config.Config, stopCh <-chan struct{}) error {
	log.Info("Starting Tencent Kubernetes Engine Auth")

	server, err := CreateServerChain(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run(stopCh)
}
