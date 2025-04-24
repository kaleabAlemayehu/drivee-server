-- name: GetCarPhotos :many
SELECT id, car_id, photo_url FROM car_photos WHERE car_id = $1 ORDER BY created_at;

-- name: GetCarPhoto :one
SELECT id, car_id, photo_url FROM car_photos WHERE id = $1 LIMIT 1;

-- name: InsertCarPhoto :one
INSERT INTO car_photos(car_id, photo_url) VALUES($1, $2) RETURNING *;

-- name: UpdateCarPhoto :one
UPDATE car_photos cp SET photo_url=$3 FROM cars c WHERE cp.car_id = c.id AND cp.id =$1 AND c.owner_id = $2 RETURNING *;
