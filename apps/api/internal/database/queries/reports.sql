-- name: CreateReport :one
-- Inserts a new report record into the 'report' table, returning the full new row.
INSERT INTO report (
  court_id,
  courts_occupied,
  groups_waiting
) VALUES (
  $1, $2, $3
) RETURNING *;