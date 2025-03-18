
-- name: ListCars :many
SELECT id, owner_id, make, model, year, license_plate, vin_number, transmission,fuel_type, mileage, location, price_per_hour, status FROM cars ORDER BY year;


-- name: GetCar :one
SELECT id, owner_id, make, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status FROM cars WHERE id = $1 LIMIT 1;


-- name: InsertCar :one
INSERT INTO cars( owner_id, make, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, ST_SetSRID(ST_MakePoint($10, $11), 4326), $12, $13 ) RETURNING *;


-- name: UpdateCar :one
UPDATE cars SET mileage = $2, location = ST_SetSRID(ST_MakePoint($3, $4), 4326), price_per_hour = $5, status = $6 WHERE id = $1 RETURNING *;


