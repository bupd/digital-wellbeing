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
WHERE updated_at >= datetime('now', '-1 hour');

-- name: ListLastDayWmClass :many
SELECT *
FROM wmclass
WHERE created_at >= DATETIME('now', '-1 day');

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
WHERE start_time >= datetime('now', '-1 hour') -- Filter for the last 24 hours
GROUP BY wm_class, wm_name
ORDER BY duration DESC;

-- name: TopActiveDurationWinLastDay :many
SELECT *
FROM wmclass
WHERE updated_at >= datetime('now', '-1 day') -- Filter for the last 24 hours
GROUP BY wm_class, wm_name
ORDER BY active_duration DESC;

-- name: TopActiveDurationWinLastHour :many
SELECT *
FROM wmclass
WHERE start_time >= datetime('now', '-1 hour') -- Filter for the last 24 hours
GROUP BY wm_class, wm_name
ORDER BY active_duration DESC;
