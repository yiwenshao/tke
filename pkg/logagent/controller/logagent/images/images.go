package images

import (
	"fmt"
	"reflect"
	"sort"

	"tkestack.io/tke/pkg/util/containerregistry"
)

const (
	// LatestVersion is latest version of addon.
	LatestVersion = "v1.0.0"
)

type Components struct {
	LogCollector containerregistry.Image
}

func (c Components) Get(name string) *containerregistry.Image {
	v := reflect.ValueOf(c)
	for i := 0; i < v.NumField(); i++ {
		v, _ := v.Field(i).Interface().(containerregistry.Image)
		if v.Name == name {
			return &v
		}
	}
	return nil
}

var versionMap = map[string]Components{
	LatestVersion: {
		LogCollector: containerregistry.Image{Name: "log-collector", Tag: "v1.1.0"},
	},
}

func List() []string {
	items := make([]string, 0, len(versionMap))
	keys := make([]string, 0, len(versionMap))
	for key := range versionMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		v := reflect.ValueOf(versionMap[key])
		for i := 0; i < v.NumField(); i++ {
			v, _ := v.Field(i).Interface().(containerregistry.Image)
			items = append(items, v.BaseName())
		}
	}

	return items
}

func Validate(version string) error {
	_, ok := versionMap[version]
	if !ok {
		return fmt.Errorf("the component version definition corresponding to version %s could not be found", version)
	}
	return nil
}

func Get(version string) Components {
	cv, ok := versionMap[version]
	if !ok {
		panic(fmt.Sprintf("the component version definition corresponding to version %s could not be found", version))
	}
	return cv
}
