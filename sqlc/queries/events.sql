-- name: ListEvents :many
SELECT * FROM events ORDER BY event_date DESC;

-- name: GetEvent :one
SELECT * FROM events WHERE id = ? LIMIT 1;

-- name: CreateEvent :one
INSERT INTO events (title, description, location, event_date)
VALUES (?, ?, ?, ?)
RETURNING *;
