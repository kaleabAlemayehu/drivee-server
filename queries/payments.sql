-- name: ListPayments :many
SELECT id, booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id FROM payments ORDER BY created_at;

-- name: GetPayment :one
SELECT id, booking_id, renter_id, owner_id, amount, payment_status, payment_method, transaction_id FROM payments WHERE id=$1;
