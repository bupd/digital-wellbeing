// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: mouse.sql

package database

import (
	"context"
)

const addMouseDown = `-- name: AddMouseDown :one
INSERT INTO mouse (event_type, button)
VALUES (?1, ?2)
RETURNING id, event_type, button, created_at
`

type AddMouseDownParams struct {
	EventType string
	Button    string
}

func (q *Queries) AddMouseDown(ctx context.Context, arg AddMouseDownParams) (Mouse, error) {
	row := q.db.QueryRowContext(ctx, addMouseDown, arg.EventType, arg.Button)
	var i Mouse
	err := row.Scan(
		&i.ID,
		&i.EventType,
		&i.Button,
		&i.CreatedAt,
	)
	return i, err
}

const listAllMouse = `-- name: ListAllMouse :many
SELECT id, event_type, button, created_at FROM mouse
`

func (q *Queries) ListAllMouse(ctx context.Context) ([]Mouse, error) {
	rows, err := q.db.QueryContext(ctx, listAllMouse)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Mouse
	for rows.Next() {
		var i Mouse
		if err := rows.Scan(
			&i.ID,
			&i.EventType,
			&i.Button,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMouseEventsLast24Hours = `-- name: ListMouseEventsLast24Hours :many
SELECT event_type, button, COUNT(*) as event_count
FROM mouse
WHERE created_at >= DATETIME('now', '-1 day')
GROUP BY event_type, button
ORDER BY event_count DESC
`

type ListMouseEventsLast24HoursRow struct {
	EventType  string
	Button     string
	EventCount int64
}

func (q *Queries) ListMouseEventsLast24Hours(ctx context.Context) ([]ListMouseEventsLast24HoursRow, error) {
	rows, err := q.db.QueryContext(ctx, listMouseEventsLast24Hours)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMouseEventsLast24HoursRow
	for rows.Next() {
		var i ListMouseEventsLast24HoursRow
		if err := rows.Scan(&i.EventType, &i.Button, &i.EventCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMouseEventsLastHour = `-- name: ListMouseEventsLastHour :many
SELECT event_type, button, COUNT(*) as event_count
FROM mouse
WHERE created_at >= DATETIME('now', '-1 hour')
GROUP BY event_type, button
ORDER BY event_count DESC
`

type ListMouseEventsLastHourRow struct {
	EventType  string
	Button     string
	EventCount int64
}

func (q *Queries) ListMouseEventsLastHour(ctx context.Context) ([]ListMouseEventsLastHourRow, error) {
	rows, err := q.db.QueryContext(ctx, listMouseEventsLastHour)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMouseEventsLastHourRow
	for rows.Next() {
		var i ListMouseEventsLastHourRow
		if err := rows.Scan(&i.EventType, &i.Button, &i.EventCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
