package starter

import (
	"strconv"
	"sync"
)

type Metrics struct {
	mu     sync.Mutex
	counts map[string]uint64
}

func New() *Metrics { return &Metrics{counts: map[string]uint64{}} }
func (m *Metrics) Record(route string, status int) {

	m.mu.Lock()
	class := status / 100
	key := route + "|" + strconv.Itoa(class) + "xx"
	m.counts[key]++
	m.mu.Unlock()

}
func (m *Metrics) Snapshot() map[string]uint64 {
	m.mu.Lock()
	result := make(map[string]uint64, len(m.counts))

	for k, v := range m.counts {
		result[k] = v
	}

	m.mu.Unlock()
	return result

}
