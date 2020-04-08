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

// ListKeys implements the interface required by DeltaFIFO to list the keys we
// already know about.
func (s *logcollectorCache) ListKeys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]string, 0, len(s.lcMap))
	for k := range s.lcMap {
		keys = append(keys, k)
	}
	return keys
}

// GetByKey returns the value stored in the lcMap under the given key
func (s *logcollectorCache) GetByKey(key string) (interface{}, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.lcMap[key]; ok {
		return v, true, nil
	}
	return nil, false, nil
}

func (s *logcollectorCache) get(logCollectorName string) (*cachedLogCollector, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	logCollector, ok := s.lcMap[logCollectorName]
	return logCollector, ok
}

func (s *logcollectorCache) Exist(logCollectorName string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.lcMap[logCollectorName]
	return ok
}

func (s *logcollectorCache) getOrCreate(logCollectorName string) *cachedLogCollector {
	s.mu.Lock()
	defer s.mu.Unlock()
	collector, ok := s.lcMap[logCollectorName]
	if !ok {
		collector = &cachedLogCollector{}
		s.lcMap[logCollectorName] = collector
	}
	return collector
}

func (s *logcollectorCache) set(logCollectorName string, collector *cachedLogCollector) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lcMap[logCollectorName] = collector
}

func (s *logcollectorCache) delete(logCollectorName string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.lcMap, logCollectorName)
}
