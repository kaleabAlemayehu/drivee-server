-- name: GetUserByEmail :one
SELECT id, first_name, middle_name, last_name, email, password FROM users WHERE email = $1 LIMIT 1;
