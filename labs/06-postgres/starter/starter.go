package starter

import "errors"

var ErrInvalidOffset = errors.New("offset must not be negative")

func ListTasksQuery(limit, offset int) (string, []any, error) {
	panic("TODO: implement ListTasksQuery")
}
