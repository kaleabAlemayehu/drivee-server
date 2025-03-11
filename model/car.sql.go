// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: car.sql

package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getCar = `-- name: GetCar :one
SELECT id, owner_id, make, model, year, license_plate, vin_number, transmission, fuel_type, mileage, location, price_per_hour, status FROM cars WHERE id = $1 LIMIT 1
`

type GetCarRow struct {
	ID           uuid.UUID      `json:"id"`
	OwnerID      uuid.UUID      `json:"owner_id"`
	Make         string         `json:"make"`
	Model        string         `json:"model"`
	Year         string         `json:"year"`
	LicensePlate string         `json:"license_plate"`
	VinNumber    string         `json:"vin_number"`
	Transmission Transmission   `json:"transmission"`
	FuelType     FuelType       `json:"fuel_type"`
	Mileage      int32          `json:"mileage"`
	Location     interface{}    `json:"location"`
	PricePerHour pgtype.Numeric `json:"price_per_hour"`
	Status       StatusType     `json:"status"`
}

func (q *Queries) GetCar(ctx context.Context, id uuid.UUID) (GetCarRow, error) {
	row := q.db.QueryRow(ctx, getCar, id)
	var i GetCarRow
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Make,
		&i.Model,
		&i.Year,
		&i.LicensePlate,
		&i.VinNumber,
		&i.Transmission,
		&i.FuelType,
		&i.Mileage,
		&i.Location,
		&i.PricePerHour,
		&i.Status,
	)
	return i, err
}

const listCars = `-- name: ListCars :many
SELECT id, owner_id, make, model, year, license_plate, vin_number, transmission,fuel_type, mileage, location, price_per_hour, status FROM cars ORDER BY year
`

type ListCarsRow struct {
	ID           uuid.UUID      `json:"id"`
	OwnerID      uuid.UUID      `json:"owner_id"`
	Make         string         `json:"make"`
	Model        string         `json:"model"`
	Year         string         `json:"year"`
	LicensePlate string         `json:"license_plate"`
	VinNumber    string         `json:"vin_number"`
	Transmission Transmission   `json:"transmission"`
	FuelType     FuelType       `json:"fuel_type"`
	Mileage      int32          `json:"mileage"`
	Location     interface{}    `json:"location"`
	PricePerHour pgtype.Numeric `json:"price_per_hour"`
	Status       StatusType     `json:"status"`
}

func (q *Queries) ListCars(ctx context.Context) ([]ListCarsRow, error) {
	rows, err := q.db.Query(ctx, listCars)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCarsRow
	for rows.Next() {
		var i ListCarsRow
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Make,
			&i.Model,
			&i.Year,
			&i.LicensePlate,
			&i.VinNumber,
			&i.Transmission,
			&i.FuelType,
			&i.Mileage,
			&i.Location,
			&i.PricePerHour,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
