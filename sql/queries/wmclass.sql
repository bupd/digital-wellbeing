-- name: ListAllWmclass :many
SELECT * FROM wmclass;

-- name: AddWmClass :exec
INSERT OR REPLACE INTO wmclass (wm_class, wm_name, start_time, end_time, duration, active_duration, is_active, updated_at)
VALUES (:wm_class, :wm_name, :start_time, :end_time, :duration, :active_duration, :is_active, :updated_at);

-- name: ListWinByWmClass :many
SELECT *
FROM wmclass
WHERE wm_class = :wm_class;

-- name: GetWinByWmName :one
SELECT *
FROM wmclass
WHERE wm_name = :wm_name;
--
-- name: ListLastHourWmClass :many
SELECT *
FROM wmclass
WHERE updated_at >= datetime('now', '-1 hour')
ORDER BY duration DESC;

-- name: ListLastDayWmClass :many
SELECT *
FROM wmclass
WHERE updated_at >= DATETIME('now', '-1 day')
ORDER BY duration DESC;

-- The below are leaving out because of the bugs in sqlc library
-- beware of bug: >sqlc generate
-- # package
-- sql/queries/wmclass.sql:1:1: duplicate query name: AddWmClass

-- name: TopDurationWinLastDay :many
SELECT *
FROM wmclass
WHERE updated_at >= datetime('now', '-1 day') -- Filter for the last 24 hours
GROUP BY wm_class, wm_name
ORDER BY duration DESC;

-- name: TopDurationWinLastHour :many
SELECT *
FROM wmclass
WHERE updated_at >= datetime('now', '-1 hour') -- Filter for the last 24 hours
GROUP BY wm_class
ORDER BY duration DESC;

-- name: TopActiveDurationWinLastDay :many
SELECT wm_class, active_duration
FROM wmclass AS w
WHERE updated_at >= datetime('now', '-1 day')
  AND active_duration = (
    SELECT MAX(active_duration)
    FROM wmclass
    WHERE wm_class = w.wm_class
      AND updated_at >= datetime('now', '-1 day')
  )
ORDER BY active_duration DESC;

-- name: TopActiveDurationWinLastHour :many
SELECT wm_class, active_duration
FROM wmclass AS w
WHERE updated_at >= datetime('now', '-1 hour')
  AND active_duration = (
    SELECT MAX(active_duration)
    FROM wmclass
    WHERE wm_class = w.wm_class
      AND updated_at >= datetime('now', '-1 hour')
  )
ORDER BY active_duration DESC;
