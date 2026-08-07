package solution_test

import (
	target "bootcamp/06-postgres/solution"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestListTasksQuery(t *testing.T) {
	query, args, err := target.ListTasksQuery(500, 3)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(query, "500") || !strings.Contains(query, "$1") {
		t.Fatalf("unsafe query %q", query)
	}
	if !reflect.DeepEqual(args, []any{100, 3}) {
		t.Fatalf("args %v", args)
	}
	if _, _, err := target.ListTasksQuery(1, -1); !errors.Is(err, target.ErrInvalidOffset) {
		t.Fatalf("got %v", err)
	}
}
