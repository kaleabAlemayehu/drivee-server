
-- name: ListUser :many
SELECT id, first_name, middle_name, last_name, email, driver_license, is_owner, is_renter, phone_number,account_number, bank_name FROM users
ORDER BY email;

-- name: GetUser :one
SELECT id, first_name, middle_name, last_name, email, password, driver_license, is_owner, is_renter, phone_number, account_number, bank_name FROM users
WHERE id = $1 LIMIT 1;

-- name: InsertUser :one
INSERT INTO users (first_name, middle_name, last_name, email, password, driver_license, phone_number, account_number, bank_name)
VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET first_name = $2, middle_name = $3, last_name = $4, email = $5, password = $6, phone_number = $7, account_number = $8, bank_name = $9, driver_license= $10, is_owner=$11, is_renter=$12 WHERE  id = $1
RETURNING *;
