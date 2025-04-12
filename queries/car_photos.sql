-- name: GetCarPhotos :one
SELECT id, car_id, photo_url FROM car_photos ORDER BY created_at;

-- name: GetCarPhoto :one
SELECT id, car_id, photo_url FROM car_photos WHERE id = $1 LIMIT 1;

-- name: InsertCarPhoto :one
INSERT INTO car_photos(car_id, photo_url) VALUES($1, $2) RETURNING *;

-- name: UpdateCarPhoto :one
UPDATE car_photos SET photo_url=$2 WHERE id =$1 RETURNING *;
