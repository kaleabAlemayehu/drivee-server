package model

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// ListBookingsForOwnerParams contains the parameters for the ListBookingsForOwner function
type ListBookingsForOwnerParams struct {
	OwnerID        uuid.UUID
	StatusFilter   *string // Use pointer for optional filters
	CarIDFilter    *uuid.NullUUID
	RenterIDFilter *uuid.NullUUID
}

// ListBookingsForOwner retrieves bookings for a car owner with optional filters
func (q *Queries) ListBookingsForOwner(ctx context.Context, params ListBookingsForOwnerParams) ([]Booking, error) {
	// Build the query with proper placeholders
	query := `
        SELECT
	b.*
        FROM
            bookings b
        JOIN
            cars c ON b.car_id = c.id
        WHERE
            c.owner_id = $1
    `

	// Create args array starting with the required ownerID
	args := []interface{}{params.OwnerID}
	paramCount := 1

	// Add conditional filters
	if params.StatusFilter != nil {
		paramCount++
		query += fmt.Sprintf(" AND b.status = $%d", paramCount)
		args = append(args, *params.StatusFilter)
	}

	if params.CarIDFilter != nil {
		paramCount++
		query += fmt.Sprintf(" AND b.car_id = $%d", paramCount)
		args = append(args, *params.CarIDFilter)
	}

	if params.RenterIDFilter != nil {
		paramCount++
		query += fmt.Sprintf(" AND b.renter_id = $%d", paramCount)
		args = append(args, *params.RenterIDFilter)
	}

	// Add ordering
	query += " ORDER BY b.created_at DESC"

	// Execute the query
	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query bookings: %w", err)
	}
	defer rows.Close()

	// Parse the results
	var bookings []Booking
	for rows.Next() {
		var booking Booking
		if err := rows.Scan(
			&booking.ID,
			&booking.CarID,
			&booking.RenterID,
			&booking.StartTime,
			&booking.EndTime,
			&booking.TotalPrice,
			&booking.Status,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan booking row: %w", err)
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate booking rows: %w", err)
	}

	return bookings, nil
}
