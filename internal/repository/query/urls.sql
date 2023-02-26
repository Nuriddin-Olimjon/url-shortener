-- name: CreateURL :one
INSERT INTO urls (
    user_id,
    original_url,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING id;


-- name: SetShortURLByID :one
UPDATE urls
SET short_uri = $1,
    expires_at = $2
WHERE id = $3
RETURNING *;


-- name: IncreaseURLRequestedCount :exec
UPDATE urls
SET requested_count = requested_count + 1
WHERE short_uri = $1;


-- name: GetURLByShortURI :one
SELECT *
FROM urls
WHERE short_uri = $1;


-- name: GetUserURLS :many
SELECT *
FROM urls
WHERE user_id = $1
ORDER BY requested_count DESC, id DESC;
