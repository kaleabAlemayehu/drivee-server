
-- name: ListBookings :many
SELECT id, car_id, renter_id, start_time, end_time, total_price, status FROM bookings ORDER BY start_time; 
