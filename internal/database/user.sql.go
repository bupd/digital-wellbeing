// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO user (name)
VALUES (?1)
RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const delteUser = `-- name: DelteUser :exec
DELETE FROM user
WHERE name = (?1)
`

func (q *Queries) DelteUser(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, delteUser, name)
	return err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, created_at, updated_at FROM user
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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
