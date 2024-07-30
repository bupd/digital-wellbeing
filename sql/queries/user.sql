-- name: ListUsers :many
SELECT * FROM user;

-- name: CreateUser :one
INSERT INTO user (name, created_at, updated_at) 
VALUES (:name, :created_at, :updated_at) 
RETURNING id, name, created_at, updated_at;


