-- name: GetUserByEmail :one
SELECT id, first_name, middle_name, last_name, email, password, profile_picture FROM users WHERE email = $1 LIMIT 1;

-- name: UpdateUserPasswordByID :exec
UPDATE users  set password = $2 WHERE id = $1 ;
