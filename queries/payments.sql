-- name: ListPayments :many
SELECT id, booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id FROM payments WHERE renter_id = $1 ORDER BY created_at;

-- name: GetPayment :one
SELECT id, booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id FROM payments WHERE id=$1 AND renter_id =$2;

-- name: InsertPayment :one
INSERT INTO payments(booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- TODO: this is going to be internal query obviously

-- name: UpdatePayment :one
UPDATE payments SET payment_status = $2  WHERE id = $1 RETURNING *;
