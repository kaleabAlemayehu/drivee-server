
-- name: ListOwner :many
SELECT id, first_name, middle_name, last_name, email, phone_number,account_number, bank_name FROM owner
ORDER BY email;



-- name: ListOwnerFull :many
SELECT * FROM owner
ORDER BY email;

