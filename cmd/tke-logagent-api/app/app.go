package app

import (
	"tkestack.io/tke/cmd/tke-logagent-api/app/config"
	"tkestack.io/tke/cmd/tke-logagent-api/app/options"
	"tkestack.io/tke/pkg/app"
	"tkestack.io/tke/pkg/util/log"
	commonapiserver "k8s.io/apiserver/pkg/server"
)

const commandDesc = `The TKE logagent provides log .`

// NewApp creates a App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions(basename)
	application := app.NewApp("Tencent Kubernetes Engine logagent",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)),
	)
	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		if err := opts.Complete(); err != nil {
			return err
		}

		cfg, err := config.CreateConfigFromOptions(basename, opts)
		if err != nil {
			return err
		}

		stopCh := commonapiserver.SetupSignalHandler()
		return Run(cfg, stopCh)
	}
}
