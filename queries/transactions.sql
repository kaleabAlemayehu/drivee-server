
-- name: ListTransactions :many
SELECT id, renter_id, owner_id, booking_id, amount, fee, payout_status FROM transactions ORDER BY created_at; 

