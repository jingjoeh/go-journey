package starter

import "sync"

type Metrics struct {
	mu     sync.Mutex
	counts map[string]uint64
}

func New() *Metrics                                { return &Metrics{counts: map[string]uint64{}} }
func (m *Metrics) Record(route string, status int) { panic("TODO: implement Metrics.Record") }
func (m *Metrics) Snapshot() map[string]uint64     { panic("TODO: implement Metrics.Snapshot") }
