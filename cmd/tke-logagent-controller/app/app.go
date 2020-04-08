package app

import (
	commonapiserver "k8s.io/apiserver/pkg/server"
	"tkestack.io/tke/cmd/tke-logagent-controller/app/config"
	"tkestack.io/tke/cmd/tke-logagent-controller/app/options"
	"tkestack.io/tke/pkg/app"
	"tkestack.io/tke/pkg/util/log"
)

const commandDesc = `The logagent controller manager is a daemon that embeds the core control loops. In
applications of robotics and automation, a control loop is a non-terminating
loop that regulates the state of the system. In TKE, a controller is a control
loop that watches the shared state of the project and cluster through the
apiserver and makes changes attempting to move the current state towards the
desired state.`

// NewApp creates a App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions(basename, KnownControllers(), ControllersDisabledByDefault.List())
	application := app.NewApp("Tencent Kubernetes Engine Logagent Controller Manager",
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

		cfg, err := config.CreateConfigFromOptions(basename, opts)
		if err != nil {
			return err
		}

		stopCh := commonapiserver.SetupSignalHandler()
		return Run(cfg, stopCh)
	}
}
