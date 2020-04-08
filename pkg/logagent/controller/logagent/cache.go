package logagent

import (
	"sync"
	v1 "tkestack.io/tke/api/logagent/v1"
)

type cachedLogCollector struct {
	// The cached state of the collector
	state *v1.LogAgent
}

type logcollectorCache struct {
	mu    sync.Mutex // protects lcMap
	lcMap map[string]*cachedLogCollector
}
