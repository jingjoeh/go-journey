package starter

import "errors"

var ErrInvalidOffset = errors.New("offset must not be negative")

func ListTasksQuery(limit, offset int) (string, []any, error) {

	if offset < 0 {
		return "", nil, ErrInvalidOffset
	}

	if limit < 1 {
		limit = 20
	} else if limit > 100 {
		limit = 100
	}

	query := `
	SELECT id, title
	FROM tasks
	LIMIT $1 OFFSET $2
`

	args := []any{limit, offset}

	return query, args, nil

}
