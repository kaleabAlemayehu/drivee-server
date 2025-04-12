
-- name: ListReviews :many
SELECT id, reviewer_id, target_id, booking_id, rating, comment FROM reviews ORDER BY updated_at;

-- name: GetReview :one
SELECT id, reviewer_id, target_id, booking_id, rating, comment FROM reviews WHERE id=$1;

-- name: InsertReview :one 
INSERT INTO reviews(reviewer_id, target_id, booking_id, rating, comment) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateReview :one
UPDATE reviews SET rating = $2, comment = $3 WHERE id = $1 RETURNING *;
