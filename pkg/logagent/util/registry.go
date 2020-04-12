package util

var (
	registryDomain    string
	registryNamespace string
)

func Init(domain string, namespace string) {
	registryDomain = domain
	registryNamespace = namespace
}

func GetRegistryDomain() string {
	return registryDomain
}

func GetRegistryNamespace() string {
	return registryNamespace
}
