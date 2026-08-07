package starter_test

import (
	target "bootcamp/12-observability/starter"
	"sync"
	"testing"
)

func TestMetrics(t *testing.T) {
	m := target.New()
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() { defer wg.Done(); m.Record("/tasks/{id}", 204) }()
	}
	wg.Wait()
	got := m.Snapshot()
	if got["/tasks/{id}|2xx"] != 20 {
		t.Fatalf("snapshot %v", got)
	}
	got["/tasks/{id}|2xx"] = 0
	if m.Snapshot()["/tasks/{id}|2xx"] != 20 {
		t.Fatal("snapshot aliases internal map")
	}
}
