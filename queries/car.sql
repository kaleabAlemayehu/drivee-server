
-- name: ListCars :many
SELECT id, owner_id, make, model, year, license_plate, vin_number, transmission,fuel_type, mileage, location, price_per_hour, status FROM cars ORDER BY year;


-- name: GetCar :one
SELECT id, owner_id, make, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status FROM cars WHERE id = $1 LIMIT 1;


