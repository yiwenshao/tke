package app

import (
	"context"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apiserver/pkg/server/healthz"
	"os"
	"time"
	"tkestack.io/tke/api/notify"
	"tkestack.io/tke/cmd/tke-logagent-controller/app/config"
	"tkestack.io/tke/pkg/controller"
	"tkestack.io/tke/pkg/util/leaderelection"
	"tkestack.io/tke/pkg/util/leaderelection/resourcelock"
	"tkestack.io/tke/pkg/util/log"
)

// Run runs the specified notify controller manager. This should never exit.
func Run(cfg *config.Config, stopCh <-chan struct{}) error {
	log.Info("Starting Tencent Kubernetes Engine logagent controller manager")

	// Setup any health checks we will want to use.
	var checks []healthz.HealthChecker
	var electionChecker *leaderelection.HealthzAdaptor
	if cfg.Component.LeaderElection.LeaderElect {
		electionChecker = leaderelection.NewLeaderHealthzAdaptor(time.Second * 20)
		checks = append(checks, electionChecker)
	}

	// Start the controller manager HTTP server
	// serverMux is the handler for these controller *after* authn/authz filters have been applied
	serverMux := controller.NewBaseHandler(&cfg.Component.Debugging, checks...)
	handler := controller.BuildHandlerChain(serverMux, &cfg.Authorization, &cfg.Authentication, notify.Codecs)
	if _, err := cfg.SecureServing.Serve(handler, 0, stopCh); err != nil {
		return err
	}

	run := func(ctx context.Context) {
		rootClientBuilder := controller.SimpleControllerClientBuilder{
			ClientConfig: cfg.LogagentAPIServerClientConfig,
		}

		controllerContext, err := CreateControllerContext(cfg, rootClientBuilder, ctx.Done())
		if err != nil {
			log.Fatalf("error building controller context: %v", err)
		}

		if err := StartControllers(controllerContext, NewControllerInitializers(), serverMux); err != nil {
			log.Fatalf("error starting controllers: %v", err)
		}

		controllerContext.InformerFactory.Start(controllerContext.Stop)
		close(controllerContext.InformersStarted)

		select {}
	}

	ctx, cancel := context.WithCancel(context.TODO())
	go func() {
		<-stopCh
		cancel()
	}()

	if !cfg.Component.LeaderElection.LeaderElect {
		run(ctx)
		panic("unreachable")
	}

	id, err := os.Hostname()
	if err != nil {
		return err
	}

	// add a uniquifier so that two processes on the same host don't accidentally both become active
	id = id + "_" + string(uuid.NewUUID())
	rl := resourcelock.NewLogagent("tke-logagent-controller",
		cfg.LeaderElectionClient.LogagentV1(),
		resourcelock.Config{
			Identity: id,
		})

	leaderelection.RunOrDie(ctx, leaderelection.ElectionConfig{
		Lock:          rl,
		LeaseDuration: cfg.Component.LeaderElection.LeaseDuration.Duration,
		RenewDeadline: cfg.Component.LeaderElection.RenewDeadline.Duration,
		RetryPeriod:   cfg.Component.LeaderElection.RetryPeriod.Duration,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: run,
			OnStoppedLeading: func() {
				log.Fatalf("leaderelection lost")
			},
		},
		WatchDog: electionChecker,
		Name:     "tke-logagent-controller",
	})
	panic("unreachable")
}
