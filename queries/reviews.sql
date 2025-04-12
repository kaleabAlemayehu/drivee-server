
-- name: ListReviews :many
SELECT id, reviewer_id, target_id, booking_id, rating, comment FROM reviews ORDER BY updated_at;

