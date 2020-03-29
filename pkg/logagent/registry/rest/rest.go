package rest

import (
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	restclient "k8s.io/client-go/rest"
	"tkestack.io/tke/api/logagent"
	v1 "tkestack.io/tke/api/logagent/v1"
	"tkestack.io/tke/pkg/apiserver/storage"
	logagentstorage "tkestack.io/tke/pkg/logagent/registry/logagent/storage"
)
// StorageProvider is a REST type for core resources storage that implement
// RestStorageProvider interface
type StorageProvider struct {
	LoopbackClientConfig *restclient.Config
	PrivilegedUsername   string
}

// Implement RESTStorageProvider
var _ storage.RESTStorageProvider = &StorageProvider{}

func (s *StorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericserver.APIGroupInfo, bool) {
	apiGroupInfo := genericserver.NewDefaultAPIGroupInfo(logagent.GroupName, logagent.Scheme, logagent.ParameterCodec, logagent.Codecs)

	if apiResourceConfigSource.VersionEnabled(v1.SchemeGroupVersion) {//what is version enabled??
		apiGroupInfo.VersionedResourcesStorageMap[v1.SchemeGroupVersion.Version] = s.v1Storage(apiResourceConfigSource, restOptionsGetter, s.LoopbackClientConfig)
	}

	return apiGroupInfo, true
}

// GroupName return the api group name
func (*StorageProvider) GroupName() string {
	return logagent.GroupName
}

func (s *StorageProvider) v1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter, loopbackClientConfig *restclient.Config) map[string]rest.Storage {
	//do we need client??
	storageMap := make(map[string]rest.Storage)
	{
		logagentRest := logagentstorage.NewStorage(restOptionsGetter, s.PrivilegedUsername)
		storageMap["logagents"] = logagentRest.LogAgent
		storageMap["logagents/status"] = logagentRest.Status
	}
	return storageMap
}
