-- name: CreateURL :one
INSERT INTO urls (
    user_id,
    requested_count,
    original_url
) VALUES (
    $1, $2, $3
) RETURNING id;


-- name: SetShortURLByID :exec
UPDATE urls
SET short_uri = $1
WHERE id = $2;


-- name: GetURLByID :one
SELECT *
FROM urls
WHERE id = $1;


-- name: GetUserURLS :many
SELECT *
FROM urls
WHERE user_id = $1;
