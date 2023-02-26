-- name: CreateUser :one
INSERT INTO users (
  username,
  full_name,
  password
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;


-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;
