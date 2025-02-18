-- name: ListAllWindows :many
SELECT * FROM windows;

-- name: AddWindows :one
INSERT OR REPLACE INTO windows (wm_class, is_active, updated_at)
VALUES (:wm_class, :is_active, :updated_at)
RETURNING *;

-- name: ListLastHourWindows :many
SELECT *
FROM windows
WHERE updated_at >= datetime('now', '-1 hour');

-- name: ListLastDayWindows :many
SELECT *
FROM windows
WHERE created_at >= DATETIME('now', '-1 day');

-- name: TopWindowsLastDay :many
SELECT wm_class, COUNT(*) as event_count
FROM windows
WHERE updated_at >= datetime('now', '-1 day') -- Filter for the last 24 hours
GROUP BY wm_class, wm_name
ORDER BY event_count DESC;

-- name: TopWindowsLastHour :many
SELECT wm_class, COUNT(*) as event_count
FROM windows
WHERE start_time >= datetime('now', '-1 hour') -- Filter for the last 24 hours
GROUP BY wm_class
ORDER BY event_count DESC;
