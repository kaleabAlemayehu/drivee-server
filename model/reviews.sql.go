// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: reviews.sql

package model

import (
	"context"

	"github.com/google/uuid"
)

const getReview = `-- name: GetReview :one
SELECT id, reviewer_id, target_id, booking_id, rating, comment FROM reviews WHERE id=$1
`

type GetReviewRow struct {
	ID         uuid.UUID `json:"id"`
	ReviewerID uuid.UUID `json:"reviewer_id"`
	TargetID   uuid.UUID `json:"target_id"`
	BookingID  uuid.UUID `json:"booking_id"`
	Rating     int32     `json:"rating"`
	Comment    string    `json:"comment"`
}

func (q *Queries) GetReview(ctx context.Context, id uuid.UUID) (GetReviewRow, error) {
	row := q.db.QueryRow(ctx, getReview, id)
	var i GetReviewRow
	err := row.Scan(
		&i.ID,
		&i.ReviewerID,
		&i.TargetID,
		&i.BookingID,
		&i.Rating,
		&i.Comment,
	)
	return i, err
}

const insertReview = `-- name: InsertReview :one
INSERT INTO reviews(reviewer_id, target_id, booking_id, rating, comment) VALUES ($1, $2, $3, $4, $5) RETURNING id, reviewer_id, target_id, booking_id, rating, comment, created_at, updated_at
`

type InsertReviewParams struct {
	ReviewerID uuid.UUID `json:"reviewer_id"`
	TargetID   uuid.UUID `json:"target_id"`
	BookingID  uuid.UUID `json:"booking_id"`
	Rating     int32     `json:"rating"`
	Comment    string    `json:"comment"`
}

func (q *Queries) InsertReview(ctx context.Context, arg InsertReviewParams) (Review, error) {
	row := q.db.QueryRow(ctx, insertReview,
		arg.ReviewerID,
		arg.TargetID,
		arg.BookingID,
		arg.Rating,
		arg.Comment,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.ReviewerID,
		&i.TargetID,
		&i.BookingID,
		&i.Rating,
		&i.Comment,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listReviews = `-- name: ListReviews :many
SELECT id, reviewer_id, target_id, booking_id, rating, comment FROM reviews ORDER BY updated_at
`

type ListReviewsRow struct {
	ID         uuid.UUID `json:"id"`
	ReviewerID uuid.UUID `json:"reviewer_id"`
	TargetID   uuid.UUID `json:"target_id"`
	BookingID  uuid.UUID `json:"booking_id"`
	Rating     int32     `json:"rating"`
	Comment    string    `json:"comment"`
}

func (q *Queries) ListReviews(ctx context.Context) ([]ListReviewsRow, error) {
	rows, err := q.db.Query(ctx, listReviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListReviewsRow
	for rows.Next() {
		var i ListReviewsRow
		if err := rows.Scan(
			&i.ID,
			&i.ReviewerID,
			&i.TargetID,
			&i.BookingID,
			&i.Rating,
			&i.Comment,
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
