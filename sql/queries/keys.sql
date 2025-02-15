-- name: ListAllKeys :many
SELECT * FROM keys;

-- name: AddKey :one
INSERT INTO keys (keyname, keycode)
VALUES (:keyname, :keycode)
RETURNING id, keyname, keycode, created_at;

-- name: ListKeysPressedLastHour :many
SELECT keyname, keycode, COUNT(*) as press_count
FROM keys
WHERE created_at >= DATETIME('now', '-1 hour')
GROUP BY keyname, keycode
ORDER BY press_count DESC;

-- name: ListKeysPressedLast24Hours :many
SELECT keyname, keycode, COUNT(*) as press_count
FROM keys
WHERE created_at >= DATETIME('now', '-1 day')
GROUP BY keyname, keycode
ORDER BY press_count DESC;
