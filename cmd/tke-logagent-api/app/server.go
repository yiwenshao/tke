package app

import (
	"tkestack.io/tke/cmd/tke-logagent-api/app/config"
	"tkestack.io/tke/pkg/logagent/apiserver"
	"tkestack.io/tke/pkg/platform/apiserver/filter"
	"tkestack.io/tke/pkg/util/log"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

// CreateServerChain creates the auth connected via delegation.
func CreateServerChain(cfg *config.Config) (*genericapiserver.GenericAPIServer, error) {
	apiServerConfig := createAPIServerConfig(cfg)
	apiServer, err := CreateAPIServer(apiServerConfig, genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	if err := registerHandler(apiServer); err != nil {
		return nil, err
	}

	apiServer.GenericAPIServer.AddPostStartHookOrDie("start-auth-api-server-informers", func(context genericapiserver.PostStartHookContext) error {
		cfg.VersionedSharedInformerFactory.Start(context.StopCh)
		return nil
	})

	return apiServer.GenericAPIServer, nil
}

// CreateAPIServer creates and wires a workable tke-auth.
func CreateAPIServer(logagentConfig *apiserver.Config, delegateAPIServer genericapiserver.DelegationTarget) (*apiserver.APIServer, error) {
	return logagentConfig.Complete().New(delegateAPIServer)
}

func createAPIServerConfig(cfg *config.Config) *apiserver.Config {
	return &apiserver.Config{
		GenericConfig: &genericapiserver.RecommendedConfig{
			Config: *cfg.GenericAPIServerConfig,
		},
		ExtraConfig: apiserver.ExtraConfig{
			ServerName:              cfg.ServerName,
			VersionedInformers:      cfg.VersionedSharedInformerFactory,
			StorageFactory:          cfg.StorageFactory,
			APIResourceConfigSource: cfg.StorageFactory.APIResourceConfigSource,
			PrivilegedUsername:      cfg.PrivilegedUsername,
			PlatformClient:          cfg.PlatformClient,
		},
	}
}

func createFilterChain(apiServer *genericapiserver.GenericAPIServer) {
	apiServer.Handler.FullHandlerChain = filter.WithCluster(apiServer.Handler.FullHandlerChain)
	apiServer.Handler.FullHandlerChain = filter.WithRequestBody(apiServer.Handler.FullHandlerChain)
	apiServer.Handler.FullHandlerChain = filter.WithFuzzyResource(apiServer.Handler.FullHandlerChain)
}

func registerHandler(apiServer *apiserver.APIServer) error {
	createFilterChain(apiServer.GenericAPIServer)
	log.Info("All of http handlers registered", log.Strings("paths", apiServer.GenericAPIServer.Handler.ListedPaths()))
	return nil
}
