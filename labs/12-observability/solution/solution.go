package solution

import (
	"fmt"
	"sync"
)

type Metrics struct {
	mu     sync.Mutex
	counts map[string]uint64
}

func New() *Metrics { return &Metrics{counts: map[string]uint64{}} }
func (m *Metrics) Record(route string, status int) {
	key := fmt.Sprintf("%s|%dxx", route, status/100)
	m.mu.Lock()
	m.counts[key]++
	m.mu.Unlock()
}
func (m *Metrics) Snapshot() map[string]uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make(map[string]uint64, len(m.counts))
	for key, value := range m.counts {
		out[key] = value
	}
	return out
}
