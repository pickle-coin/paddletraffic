-- name: CreateCourt :one
WITH new_location AS (
  INSERT INTO location (address_line, country_code, timezone, lat, lon, region, postal_code, place_id)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING *
), new_court AS (
  INSERT INTO court (location_id, name, court_count)
  VALUES ((SELECT id FROM new_location), $9, $10)
  RETURNING *
)
SELECT
  c.id as court_id,
  c.location_id,
  c.name as court_name,
  c.court_count,
  c.created_at as court_created_at,
  c.updated_at as court_updated_at,
  l.id as location_id_val,
  l.address_line,
  l.country_code,
  l.timezone,
  l.lat,
  l.lon,
  l.region,
  l.postal_code,
  l.place_id,
  l.created_at as location_created_at,
  l.updated_at as location_updated_at
FROM new_court c
CROSS JOIN new_location l;

-- name: GetAllCourts :many
SELECT
  c.id as court_id,
  c.location_id,
  c.name as court_name,
  c.court_count,
  c.created_at as court_created_at,
  c.updated_at as court_updated_at,
  l.id as location_id_val,
  l.address_line,
  l.country_code,
  l.timezone,
  l.lat,
  l.lon,
  l.region,
  l.postal_code,
  l.place_id,
  l.created_at as location_created_at,
  l.updated_at as location_updated_at
FROM court c
INNER JOIN location l ON c.location_id = l.id
ORDER BY c.created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountCourts :one
SELECT COUNT(*) FROM court;
