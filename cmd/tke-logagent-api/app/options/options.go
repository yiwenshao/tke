package options

import (
	"github.com/spf13/pflag"
	genericapiserveroptions "k8s.io/apiserver/pkg/server/options"
	apiserveroptions "tkestack.io/tke/pkg/apiserver/options"
	storageoptions "tkestack.io/tke/pkg/apiserver/storage/options"
	controlleroptions "tkestack.io/tke/pkg/controller/options"
	"tkestack.io/tke/pkg/util/cachesize"
	"tkestack.io/tke/pkg/util/log"
)

// Options is the main context object for the TKE monitor.
type Options struct {
	Log               *log.Options
	SecureServing     *apiserveroptions.SecureServingOptions
	Debug             *apiserveroptions.DebugOptions
	ETCD              *storageoptions.ETCDStorageOptions
	Generic           *apiserveroptions.GenericOptions
	Authentication    *apiserveroptions.AuthenticationWithAPIOptions
	Authorization     *apiserveroptions.AuthorizationOptions
	PlatformAPIClient *controlleroptions.APIServerClientOptions
	// The Registry will load its initial configuration from this file.
	// The path may be absolute or relative; relative paths are under the Monitor's current working directory.
	LogagentConfig string
}

// NewOptions creates a new Options with a default config.
func NewOptions(serverName string) *Options {
	return &Options{
		Log:            log.NewOptions(),
		SecureServing:  apiserveroptions.NewSecureServingOptions(serverName, 9999),
		Debug:          apiserveroptions.NewDebugOptions(),
		Generic:        apiserveroptions.NewGenericOptions(),
		Authentication: apiserveroptions.NewAuthenticationWithAPIOptions(),
		Authorization:  apiserveroptions.NewAuthorizationOptions(),
		ETCD:           storageoptions.NewETCDStorageOptions("/tke/logagent-api"),
		PlatformAPIClient: controlleroptions.NewAPIServerClientOptions("platform", true),
		//Auth:           NewAuthOptions(), //options/logagent.go is not used currently
	}
}

// AddFlags adds flags for a specific server to the specified FlagSet object.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	o.Log.AddFlags(fs)
	o.SecureServing.AddFlags(fs)
	o.Debug.AddFlags(fs)
	o.ETCD.AddFlags(fs)
	o.Generic.AddFlags(fs)
	o.Authentication.AddFlags(fs)
	o.Authorization.AddFlags(fs)
	o.PlatformAPIClient.AddFlags(fs)//read platform config ang generate platform client
	//reference monitor to add more flags here
}

// ApplyFlags parsing parameters from the command line or configuration file
// to the options instance.
func (o *Options) ApplyFlags() []error {
	var errs []error

	errs = append(errs, o.Log.ApplyFlags()...)
	errs = append(errs, o.SecureServing.ApplyFlags()...)
	errs = append(errs, o.Debug.ApplyFlags()...)
	errs = append(errs, o.ETCD.ApplyFlags()...)
	errs = append(errs, o.Generic.ApplyFlags()...)
	errs = append(errs, o.Authentication.ApplyFlags()...)
	errs = append(errs, o.Authorization.ApplyFlags()...)
	errs = append(errs, o.PlatformAPIClient.ApplyFlags()...)
	//errs = append(errs, o.Auth.ApplyFlags()...)

	return errs
}


// Complete set default Options.
// Should be called after tke-logagent flags parsed.
func (o *Options) Complete() error {
	if err := apiserveroptions.CompleteGenericAndSecureOptions(o.Generic, o.SecureServing); err != nil {
		return err
	}

	if o.ETCD.EnableWatchCache {
		log.Infof("Initializing cache sizes based on %dMB limit", o.Generic.TargetRAMMB)
		sizes := cachesize.NewHeuristicWatchCacheSizes(o.Generic.TargetRAMMB)
		if userSpecified, err := genericapiserveroptions.ParseWatchCacheSizes(o.ETCD.WatchCacheSizes); err == nil {
			for resource, size := range userSpecified {
				sizes[resource] = size
			}
		}

		watchCacheSizes, err := genericapiserveroptions.WriteWatchCacheSizes(sizes)
		if err != nil {
			return err
		}
		o.ETCD.WatchCacheSizes = watchCacheSizes
	}
	return nil
}
