
-- name: ListTransactions :many
SELECT id, renter_id, owner_id, booking_id, amount, fee, payout_status FROM transactions ORDER BY created_at; 

-- name: GetTransaction :one
SELECT renter_id, owner_id, booking_id, amount, fee, payout_status FROM transactions WHERE id=$1; 

-- name: InsertTransaction :one
INSERT INTO transactions(renter_id, owner_id, booking_id, amount, fee, payout_status ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateTransaction :one
UPDATE transactions SET payout_status = $2 WHERE id = $1 RETURNING *;
