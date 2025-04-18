// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: wmclass.sql

package database

import (
	"context"
	"time"
)

const addWmClass = `-- name: AddWmClass :exec
INSERT OR REPLACE INTO wmclass (wm_class, wm_name, start_time, end_time, duration, active_duration, is_active, updated_at)
VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8)
`

type AddWmClassParams struct {
	WmClass        string
	WmName         string
	StartTime      time.Time
	EndTime        time.Time
	Duration       int64
	ActiveDuration int64
	IsActive       int64
	UpdatedAt      time.Time
}

func (q *Queries) AddWmClass(ctx context.Context, arg AddWmClassParams) error {
	_, err := q.db.ExecContext(ctx, addWmClass,
		arg.WmClass,
		arg.WmName,
		arg.StartTime,
		arg.EndTime,
		arg.Duration,
		arg.ActiveDuration,
		arg.IsActive,
		arg.UpdatedAt,
	)
	return err
}

const getWinByWmName = `-- name: GetWinByWmName :one
SELECT id, wm_class, wm_name, start_time, end_time, duration, active_duration, is_active, created_at, updated_at
FROM wmclass
WHERE wm_name = ?1
`

func (q *Queries) GetWinByWmName(ctx context.Context, wmName string) (Wmclass, error) {
	row := q.db.QueryRowContext(ctx, getWinByWmName, wmName)
	var i Wmclass
	err := row.Scan(
		&i.ID,
		&i.WmClass,
		&i.WmName,
		&i.StartTime,
		&i.EndTime,
		&i.Duration,
		&i.ActiveDuration,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllWmclass = `-- name: ListAllWmclass :many
SELECT id, wm_class, wm_name, start_time, end_time, duration, active_duration, is_active, created_at, updated_at FROM wmclass
`

func (q *Queries) ListAllWmclass(ctx context.Context) ([]Wmclass, error) {
	rows, err := q.db.QueryContext(ctx, listAllWmclass)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Wmclass
	for rows.Next() {
		var i Wmclass
		if err := rows.Scan(
			&i.ID,
			&i.WmClass,
			&i.WmName,
			&i.StartTime,
			&i.EndTime,
			&i.Duration,
			&i.ActiveDuration,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listLastDayWmClass = `-- name: ListLastDayWmClass :many
SELECT id, wm_class, wm_name, start_time, end_time, duration, active_duration, is_active, created_at, updated_at
FROM wmclass
WHERE created_at >= DATETIME('now', '-1 day')
`

func (q *Queries) ListLastDayWmClass(ctx context.Context) ([]Wmclass, error) {
	rows, err := q.db.QueryContext(ctx, listLastDayWmClass)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Wmclass
	for rows.Next() {
		var i Wmclass
		if err := rows.Scan(
			&i.ID,
			&i.WmClass,
			&i.WmName,
			&i.StartTime,
			&i.EndTime,
			&i.Duration,
			&i.ActiveDuration,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listLastHourWmClass = `-- name: ListLastHourWmClass :many
SELECT id, wm_class, wm_name, start_time, end_time, duration, active_duration, is_active, created_at, updated_at
FROM wmclass
WHERE updated_at >= datetime('now', '-1 hour')
`

func (q *Queries) ListLastHourWmClass(ctx context.Context) ([]Wmclass, error) {
	rows, err := q.db.QueryContext(ctx, listLastHourWmClass)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Wmclass
	for rows.Next() {
		var i Wmclass
		if err := rows.Scan(
			&i.ID,
			&i.WmClass,
			&i.WmName,
			&i.StartTime,
			&i.EndTime,
			&i.Duration,
			&i.ActiveDuration,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listWinByWmClass = `-- name: ListWinByWmClass :many
SELECT id, wm_class, wm_name, start_time, end_time, duration, active_duration, is_active, created_at, updated_at
FROM wmclass
WHERE wm_class = ?1
`

func (q *Queries) ListWinByWmClass(ctx context.Context, wmClass string) ([]Wmclass, error) {
	rows, err := q.db.QueryContext(ctx, listWinByWmClass, wmClass)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Wmclass
	for rows.Next() {
		var i Wmclass
		if err := rows.Scan(
			&i.ID,
			&i.WmClass,
			&i.WmName,
			&i.StartTime,
			&i.EndTime,
			&i.Duration,
			&i.ActiveDuration,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const topWinLastDay = `-- name: TopWinLastDay :many

SELECT wm_class, wm_name, COUNT(*) as event_count
FROM wmclass
WHERE updated_at >= datetime('now', '-1 day') -- Filter for the last 24 hours
GROUP BY wm_class, wm_name
ORDER BY event_count DESC
`

type TopWinLastDayRow struct {
	WmClass    string
	WmName     string
	EventCount int64
}

// The below are leaving out because of the bugs in sqlc library
// beware of bug: >sqlc generate
// # package
// sql/queries/wmclass.sql:1:1: duplicate query name: AddWmClass
func (q *Queries) TopWinLastDay(ctx context.Context) ([]TopWinLastDayRow, error) {
	rows, err := q.db.QueryContext(ctx, topWinLastDay)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TopWinLastDayRow
	for rows.Next() {
		var i TopWinLastDayRow
		if err := rows.Scan(&i.WmClass, &i.WmName, &i.EventCount); err != nil {
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

const topWinLastHour = `-- name: TopWinLastHour :many
SELECT wm_class, wm_name, COUNT(*) as event_count
FROM wmclass
WHERE start_time >= datetime('now', '-1 hour') -- Filter for the last 24 hours
GROUP BY wm_class
ORDER BY event_count DESC
`

type TopWinLastHourRow struct {
	WmClass    string
	WmName     string
	EventCount int64
}

func (q *Queries) TopWinLastHour(ctx context.Context) ([]TopWinLastHourRow, error) {
	rows, err := q.db.QueryContext(ctx, topWinLastHour)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TopWinLastHourRow
	for rows.Next() {
		var i TopWinLastHourRow
		if err := rows.Scan(&i.WmClass, &i.WmName, &i.EventCount); err != nil {
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
