-- name: GetBookingForOwner :one
 SELECT b.id, b.car_id, b.renter_id, b.start_time, b.end_time, b.total_price, b.status, b.booking_no FROM bookings b JOIN cars c ON b.car_id = c.id WHERE c.owner_id = $2 AND b.id=$1;

-- name: UpdateBookingForOwner :one
UPDATE bookings b SET status = $3  FROM cars c WHERE b.car_id = c.id AND c.owner_id = $2 AND b.id = $1 RETURNING b.*;

-- name: ListCarsForOwner :many
SELECT id, owner_id, thumbnail_picture, description, make, model, year, license_plate, vin_number, transmission,fuel_type, mileage, location, price_per_hour, status FROM cars WHERE owner_id = $1 ORDER BY year;

-- name: GetCarForOwner :many
SELECT id, owner_id, thumbnail_picture,make, description, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status FROM cars WHERE id = $1 AND owner_id=$2 LIMIT 1;

