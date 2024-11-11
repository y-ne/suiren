-- name: CreateUser :one
INSERT INTO users (
    username,
    password
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, username, created_at, updated_at FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users 
SET 
    username = COALESCE($2, username),
    password = COALESCE($3, password),
    updated_at = CLOCK_TIMESTAMP()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;