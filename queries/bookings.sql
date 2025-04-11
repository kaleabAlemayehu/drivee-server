
-- name: ListBookings :many
SELECT id, car_id, renter_id, start_time, end_time, total_price, status FROM bookings ORDER BY start_time; 

-- name: GetBooking :one
SELECT id, car_id, renter_id, start_time, end_time, total_price, status FROM bookings WHERE id=$1; 

-- name: InsertBooking :one
INSERT INTO bookings( car_id, renter_id, start_time, end_time, total_price, status ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;
