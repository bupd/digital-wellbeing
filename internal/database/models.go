// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"
)

type User struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
