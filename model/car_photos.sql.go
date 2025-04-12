// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: car_photos.sql

package model

import (
	"context"

	"github.com/google/uuid"
)

const getCarPhoto = `-- name: GetCarPhoto :one
SELECT id, car_id, photo_url FROM car_photos WHERE id = $1 LIMIT 1
`

type GetCarPhotoRow struct {
	ID       uuid.UUID `json:"id"`
	CarID    uuid.UUID `json:"car_id"`
	PhotoUrl string    `json:"photo_url"`
}

func (q *Queries) GetCarPhoto(ctx context.Context, id uuid.UUID) (GetCarPhotoRow, error) {
	row := q.db.QueryRow(ctx, getCarPhoto, id)
	var i GetCarPhotoRow
	err := row.Scan(&i.ID, &i.CarID, &i.PhotoUrl)
	return i, err
}

const getCarPhotos = `-- name: GetCarPhotos :many
SELECT id, car_id, photo_url FROM car_photos ORDER BY created_at
`

type GetCarPhotosRow struct {
	ID       uuid.UUID `json:"id"`
	CarID    uuid.UUID `json:"car_id"`
	PhotoUrl string    `json:"photo_url"`
}

func (q *Queries) GetCarPhotos(ctx context.Context) ([]GetCarPhotosRow, error) {
	rows, err := q.db.Query(ctx, getCarPhotos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCarPhotosRow
	for rows.Next() {
		var i GetCarPhotosRow
		if err := rows.Scan(&i.ID, &i.CarID, &i.PhotoUrl); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertCarPhoto = `-- name: InsertCarPhoto :one
INSERT INTO car_photos(car_id, photo_url) VALUES($1, $2) RETURNING id, car_id, photo_url, created_at, updated_at
`

type InsertCarPhotoParams struct {
	CarID    uuid.UUID `json:"car_id"`
	PhotoUrl string    `json:"photo_url"`
}

func (q *Queries) InsertCarPhoto(ctx context.Context, arg InsertCarPhotoParams) (CarPhoto, error) {
	row := q.db.QueryRow(ctx, insertCarPhoto, arg.CarID, arg.PhotoUrl)
	var i CarPhoto
	err := row.Scan(
		&i.ID,
		&i.CarID,
		&i.PhotoUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCarPhoto = `-- name: UpdateCarPhoto :one
UPDATE car_photos SET photo_url=$2 WHERE id =$1 RETURNING id, car_id, photo_url, created_at, updated_at
`

type UpdateCarPhotoParams struct {
	ID       uuid.UUID `json:"id"`
	PhotoUrl string    `json:"photo_url"`
}

func (q *Queries) UpdateCarPhoto(ctx context.Context, arg UpdateCarPhotoParams) (CarPhoto, error) {
	row := q.db.QueryRow(ctx, updateCarPhoto, arg.ID, arg.PhotoUrl)
	var i CarPhoto
	err := row.Scan(
		&i.ID,
		&i.CarID,
		&i.PhotoUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
