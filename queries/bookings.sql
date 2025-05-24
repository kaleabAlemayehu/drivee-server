
-- name: ListBookingsForRenter :many
SELECT id, car_id, renter_id, start_time, end_time, total_price, status FROM bookings WHERE renter_id=$1 AND (status = sqlc.narg('status') OR sqlc.narg('status') IS NULL) ORDER BY start_time; 

-- name: GetBookingForRenter :one
SELECT id, car_id, renter_id, start_time, end_time, total_price, status FROM bookings WHERE id=$1 AND renter_id =$2; 


-- name: InsertBooking :one
INSERT INTO bookings( car_id, renter_id, start_time, end_time, total_price) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateBookingForRenter :one
UPDATE bookings SET start_time = $3, end_time = $4, total_price = $5 , status='pending' WHERE id = $1 AND renter_id = $2 RETURNING *;
