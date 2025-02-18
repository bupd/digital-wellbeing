-- name: ListAllWmclass :many
SELECT * FROM wmclass;

-- name: AddWmclass :one
INSERT OR REPLACE INTO wmclass (
  wm_class,
  wm_name,
  start_time,
  end_time,
  duration,
  total_count,
  is_active,
  updated_at
)
VALUES (
  :wm_class,
  :wm_name,
  :start_time,
  :end_time,
  :duration,
  :total_count,
  :is_active,
  :updated_at
)
RETURNING *;

-- name: ListWinByWmClass :many
SELECT *
FROM wmclass
WHERE wm_class = :wm_class;

-- name: ListWinByWmName :one
SELECT *
FROM wmclass
WHERE wm_name = :wm_name;

-- name: ListLastHourWindows :many
SELECT *
FROM wmclass
WHERE updated_at >= datetime('now', '-1 hour');

-- name: ListLastDayWindows :many
SELECT *
FROM wmclass
WHERE created_at >= DATETIME('now', '-1 day')

-- name: TopWinLastDay :many
SELECT wm_class, wm_name, COUNT(*) as event_count
FROM wmclass
WHERE updated_at >= datetime('now', '-1 day') -- Filter for the last 24 hours
GROUP BY wm_class, wm_name
ORDER BY event_count DESC;

-- name: TopWinLastHour :many
SELECT wm_class, wm_name, COUNT(*) as event_count
FROM wmclass
WHERE start_time >= datetime('now', '-1 hour') -- Filter for the last 24 hours
GROUP BY wm_class
ORDER BY event_count DESC;
