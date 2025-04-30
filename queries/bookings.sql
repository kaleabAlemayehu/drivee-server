
-- name: ListBookingsForRenter :many
SELECT id, car_id, renter_id, start_time, end_time, total_price, status FROM bookings WHERE renter_id=$1 ORDER BY start_time; 

-- name: ListBookingsForOwner :many
SELECT b.id AS booking_id, b.car_id, b.renter_id, b.start_time, b.end_time, b.total_price, b.status 
	FROM bookings b JOIN cars c ON b.car_id = c.id
	JOIN users u ON b.renter_id = u.id WHERE c.owner_id = $1;

-- name: GetBookingForRenter :one
SELECT id, car_id, renter_id, start_time, end_time, total_price, status FROM bookings WHERE id=$1 AND renter_id =$2; 

-- name: GetBookingForOwner :one
SELECT b.id AS booking_id, b.car_id, b.renter_id, b.start_time, b.end_time, b.total_price, b.status 
	FROM bookings b JOIN cars c ON b.car_id = c.id
	JOIN users u ON b.renter_id = u.id WHERE c.owner_id = $2 AND b.id=$1;

-- name: InsertBooking :one
INSERT INTO bookings( car_id, renter_id, start_time, end_time, total_price) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateBookingForOwner :one
UPDATE bookings b SET status = $3  FROM cars c WHERE b.car_id = c.id AND c.owner_id = $2 AND b.id = $1 RETURNING b.*;

-- name: UpdateBookingForRenter :one
UPDATE bookings SET start_time = $3, end_time = $4, total_price = $5 , status='pending' WHERE id = $1 AND renter_id = $2 RETURNING *;
