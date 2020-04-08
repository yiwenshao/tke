package options

import (
	"github.com/spf13/pflag"
	apiserveroptions "tkestack.io/tke/pkg/apiserver/options"
	controlleroptions "tkestack.io/tke/pkg/controller/options"
	"tkestack.io/tke/pkg/util/log"
)

// Options is the main context object for the TKE controller manager.
type Options struct {
	Log               *log.Options
	Debug             *apiserveroptions.DebugOptions
	SecureServing     *apiserveroptions.SecureServingOptions
	Component         *controlleroptions.ComponentOptions
	LogagentAPIClient  *controlleroptions.APIServerClientOptions
	// The Registry will load its initial configuration from this file.
	// The path may be absolute or relative; relative paths are under the Monitor's current working directory.
	LogagentConfig string
}

// NewOptions creates a new Options with a default config.
func NewOptions(serverName string, allControllers []string, disabledByDefaultControllers []string) *Options {
	return &Options{
		Log:             log.NewOptions(),
		Debug:           apiserveroptions.NewDebugOptions(),
		SecureServing:   apiserveroptions.NewSecureServingOptions(serverName, 9998),
		Component:       controlleroptions.NewComponentOptions(allControllers, disabledByDefaultControllers),
		LogagentAPIClient: controlleroptions.NewAPIServerClientOptions("logagent", true),
	}
}

// AddFlags adds flags for a specific server to the specified FlagSet object.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	o.Log.AddFlags(fs)
	o.Debug.AddFlags(fs)
	o.SecureServing.AddFlags(fs)
	o.Component.AddFlags(fs)
	o.LogagentAPIClient.AddFlags(fs)
}

// ApplyFlags parsing parameters from the command line or configuration file
// to the options instance.
func (o *Options) ApplyFlags() []error {
	var errs []error

	errs = append(errs, o.Log.ApplyFlags()...)
	errs = append(errs, o.Debug.ApplyFlags()...)
	errs = append(errs, o.SecureServing.ApplyFlags()...)
	errs = append(errs, o.Component.ApplyFlags()...)
	errs = append(errs, o.LogagentAPIClient.ApplyFlags()...)

	return errs
}


