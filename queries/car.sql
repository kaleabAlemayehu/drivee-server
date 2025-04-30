-- TODO: will add pagenation using LIMIT and OFFSET

-- name: ListCars :many
SELECT id, owner_id, thumbnail_picture, description, make, model, year, license_plate, vin_number, transmission,fuel_type, mileage, location, price_per_hour, status FROM cars ORDER BY year;

-- name: GetCar :one
SELECT id, owner_id, thumbnail_picture,make, description, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status FROM cars WHERE id = $1 LIMIT 1;


-- name: InsertCar :one
INSERT INTO cars( owner_id, make, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status, thumbnail_picture,description ) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, ST_SetSRID(ST_MakePoint($10, $11), 4326), $12, $13, $14 , $15) RETURNING *;


-- name: UpdateCar :one
UPDATE cars SET mileage = $3, location = ST_SetSRID(ST_MakePoint($4, $5), 4326), price_per_hour = $6, status = $7 WHERE id = $1 AND owner_id = $2 RETURNING *;


