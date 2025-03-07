
-- name: ListOwner :many
SELECT id, first_name, middle_name, last_name, email, phone_number,account_number, bank_name FROM owner
ORDER BY email;

-- name: GetOwner :one
SELECT id, first_name, middle_name, last_name, email, phone_number, account_number, bank_name FROM owner
WHERE id = $1 LIMIT 1;
