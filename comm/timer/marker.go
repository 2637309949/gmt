package timer

import (
	"comm/logger"
	"time"
)

// Marker defined TODO
type Marker struct {
	tm map[string]time.Time
}

// NewMark defiend TODO
func NewMark()  *Marker {
	return &Marker{tm: map[string]time.Time{}}
}

// Marker defined TODO
func (m *Marker) Mark(name string) {
	m.tm[name] = time.Now()
}

// Init defined TODO
func (m *Marker) Init(name string) {
	logger.Infof("Starting %v timer", name)
	for k := range m.tm {
		logger.Infof("%v latency %v", k, time.Since(m.tm[k]).String())
	}
}
