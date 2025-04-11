-- name: ListPayments :many
SELECT id, booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id FROM payments ORDER BY created_at;

-- name: GetPayment :one
SELECT id, booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id FROM payments WHERE id=$1;

-- name: InsertPayment :one
INSERT INTO payments(booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

