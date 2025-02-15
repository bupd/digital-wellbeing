-- name: ListAllMouse :many
SELECT * FROM mouse;

-- name: AddMouseDown :one
INSERT INTO mouse (event_type, button)
VALUES (:event_type, :button)
RETURNING id, event_type, button, created_at;

-- name: ListMouseEventsLastHour :many
SELECT event_type, button, COUNT(*) as event_count
FROM mouse
WHERE created_at >= DATETIME('now', '-1 hour')
GROUP BY event_type, button
ORDER BY event_count DESC;

-- name: ListMouseEventsLast24Hours :many
SELECT event_type, button, COUNT(*) as event_count
FROM mouse
WHERE created_at >= DATETIME('now', '-1 day')
GROUP BY event_type, button
ORDER BY event_count DESC;
