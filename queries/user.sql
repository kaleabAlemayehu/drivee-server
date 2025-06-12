
-- name: ListUser :many
SELECT id, first_name, middle_name, last_name, email, driver_license, is_owner, is_renter, phone_number,account_number, bank_name, profile_picture FROM users
ORDER BY email;

-- TODO: adding route and handler that can use getme
-- name: GetMe :one
SELECT id, first_name, middle_name, last_name, email, password, driver_license, is_owner, is_renter, phone_number, account_number, bank_name, profile_picture  FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT id, first_name, middle_name, last_name, email,  profile_picture  FROM users
WHERE id = $1 LIMIT 1;

-- name: InsertUser :one
INSERT INTO users (first_name, email, password, profile_picture) VALUES ( $1, $2, $3, $4) RETURNING *;

-- name: InsertUserSSO :one
INSERT INTO users (first_name, last_name, email, profile_picture) VALUES ( $1, $2, $3, $4 ) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET first_name = $2, middle_name = $3, last_name = $4, email = $5, password = $6, phone_number = $7, account_number = $8, bank_name = $9, driver_license= $10, is_owner=$11, is_renter=$12 , profile_picture=$13 WHERE  id = $1
RETURNING *;


