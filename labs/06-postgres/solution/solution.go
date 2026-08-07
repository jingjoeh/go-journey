package solution

import "errors"

var ErrInvalidOffset = errors.New("offset must not be negative")

func ListTasksQuery(limit, offset int) (string, []any, error) {
	if offset < 0 {
		return "", nil, ErrInvalidOffset
	}
	if limit < 1 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	const query = "SELECT id, title FROM tasks ORDER BY id LIMIT $1 OFFSET $2"
	return query, []any{limit, offset}, nil
}
