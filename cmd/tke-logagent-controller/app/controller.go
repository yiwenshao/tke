package app

import (
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/server/mux"
	"net/http"
	"time"
	"tkestack.io/tke/pkg/util/log"
)

const (
	// ControllerStartJitter used when starting controller managers.
	ControllerStartJitter = 1.0
)

// ControllersDisabledByDefault configured all controllers that are turned off
// by default.
var ControllersDisabledByDefault = sets.NewString()

// KnownControllers returns the known controllers.
func KnownControllers() []string {
	ret := sets.StringKeySet(NewControllerInitializers())
	return ret.List()
}




// NewControllerInitializers is a public map of named controller groups (you can start more than one in an init func)
// paired to their InitFunc.  This allows for structured downstream composition and subdivision.
func NewControllerInitializers() map[string]InitFunc {
	controllers := map[string]InitFunc{}

	controllers["logagent"] = startLogagentController
	return controllers
}


// StartControllers to start the controller.
func StartControllers(ctx ControllerContext, controllers map[string]InitFunc, unsecuredMux *mux.PathRecorderMux) error {
	for controllerName, initFn := range controllers {
		if !ctx.IsControllerEnabled(controllerName) {
			log.Warnf("%q is disabled", controllerName)
			continue
		}
		time.Sleep(wait.Jitter(ctx.ControllerStartInterval, ControllerStartJitter))
		log.Infof("Starting %q", controllerName)
		debugHandler, started, err := initFn(ctx)
		if err != nil {
			log.Errorf("Error starting %q", controllerName)
			return err
		}
		if !started {
			log.Warnf("Skipping %q", controllerName)
			continue
		}
		if debugHandler != nil && unsecuredMux != nil {
			basePath := "/debug/controllers/" + controllerName
			unsecuredMux.UnlistedHandle(basePath, http.StripPrefix(basePath, debugHandler))
			unsecuredMux.UnlistedHandlePrefix(basePath+"/", http.StripPrefix(basePath, debugHandler))
		}
		log.Infof("Started %q", controllerName)
	}

	return nil
}
