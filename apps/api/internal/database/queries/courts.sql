-- name: CreateCourt :one
WITH new_location AS (
  INSERT INTO location (address_line, country_code, timezone, lat, lon, region, postal_code, place_id)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING id
)
INSERT INTO court (location_id, name, court_count)
VALUES ((SELECT id FROM new_location), $9, $10)
RETURNING *;

-- name: GetAllCourts :many
SELECT * FROM court
ORDER BY created_at DESC;
