package options

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagDomain           = "domain"
	flagNamespace        = "namespace"
)

const (
	configDomain           = "features.domain"
	configNamespace        = "features.namespace"
)

type FeatureOptions struct {
	Domain           string
	Namespace        string
}

func NewFeatureOptions() *FeatureOptions {
	return &FeatureOptions{

	}
}

// AddFlags adds flags for console to the specified FlagSet object.
func (o *FeatureOptions) AddFlags(fs *pflag.FlagSet) {
	fs.String(flagDomain, o.Domain,"registry domain")
	_ = viper.BindPFlag(configDomain,fs.Lookup(flagDomain))

	fs.String(flagNamespace, o.Namespace, "registry namespace")
	_ = viper.BindPFlag(configNamespace, fs.Lookup(flagNamespace))

}


// ApplyFlags parsing parameters from the command line or configuration file
// to the options instance.
func (o *FeatureOptions) ApplyFlags() []error {
	var errs []error
	o.Domain = viper.GetString(configDomain)
	o.Namespace = viper.GetString(configNamespace)

	if len(o.Domain) == 0 || len(o.Namespace) == 0 {
		errs = append(errs, fmt.Errorf("%s and %s must be specified", configDomain, configNamespace))
	}
	return errs
}
