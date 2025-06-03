-- name: InsertToken :one
INSERT INTO token(token ,expires_at ,user_id ) VAlUES ($1, $2, $3) RETURNING *;

-- name: GetToken :one
SELECT * FROM token WHERE token = $1 AND expires_at > $2;


