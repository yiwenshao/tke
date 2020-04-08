package app

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
	"time"
	v1 "tkestack.io/tke/api/logagent/v1"
	controlleragent "tkestack.io/tke/pkg/logagent/controller/logagent"

)

const (
	channelSyncPeriod      = 30 * time.Second
	concurrentChannelSyncs = 10

	messageRequestSyncPeriod      = 5 * time.Minute
	concurrentMessageRequestSyncs = 10
)

func startLogagentController(ctx ControllerContext) (http.Handler, bool, error) {
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: v1.GroupName, Version: "v1", Resource: "logagents"}] {
		return nil, false, nil
	}
	ctrl := controlleragent.NewController(ctx.ClientBuilder.ClientOrDie("logagent-controller"),ctx.InformerFactory.Logagent().V1().LogAgents(),channelSyncPeriod)
	go ctrl.Run(concurrentChannelSyncs, ctx.Stop)
	return nil, true, nil
}
