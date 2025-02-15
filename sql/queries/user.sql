-- name: ListUsers :many
SELECT * FROM user;

-- name: CreateUser :one
INSERT INTO user (name)
VALUES (:name)
RETURNING id, name, created_at, updated_at;

-- name: DelteUser :exec
DELETE FROM user
WHERE name = (:name);
