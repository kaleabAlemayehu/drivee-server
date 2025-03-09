
-- name: ListOwner :many
SELECT id, first_name, middle_name, last_name, email, phone_number,account_number, bank_name FROM owner
ORDER BY email;

-- name: GetOwner :one
SELECT id, first_name, middle_name, last_name, email, password, phone_number, account_number, bank_name FROM owner
WHERE id = $1 LIMIT 1;

-- name: InsertOwner :one
INSERT INTO owner (first_name, middle_name, last_name, email, password, phone_number, account_number, bank_name)
VALUES ( $1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdatedOwner :one
UPDATE owner SET first_name = $2, middle_name = $3, last_name = $4, email = $5, password = $6, phone_number = $7, account_number = $8, bank_name = $9 WHERE  id = $1
RETURNING *;
