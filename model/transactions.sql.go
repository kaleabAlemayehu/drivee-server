// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: transactions.sql

package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getTransaction = `-- name: GetTransaction :one
SELECT renter_id, owner_id, booking_id, amount, fee, payout_status FROM transactions WHERE id=$1
`

type GetTransactionRow struct {
	RenterID     uuid.UUID      `json:"renter_id"`
	OwnerID      uuid.UUID      `json:"owner_id"`
	BookingID    uuid.UUID      `json:"booking_id"`
	Amount       pgtype.Numeric `json:"amount"`
	Fee          pgtype.Numeric `json:"fee"`
	PayoutStatus PayoutStatus   `json:"payout_status"`
}

func (q *Queries) GetTransaction(ctx context.Context, id uuid.UUID) (GetTransactionRow, error) {
	row := q.db.QueryRow(ctx, getTransaction, id)
	var i GetTransactionRow
	err := row.Scan(
		&i.RenterID,
		&i.OwnerID,
		&i.BookingID,
		&i.Amount,
		&i.Fee,
		&i.PayoutStatus,
	)
	return i, err
}

const insertTransaction = `-- name: InsertTransaction :one
INSERT INTO transactions(renter_id, owner_id, booking_id, amount, fee, payout_status ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, renter_id, owner_id, booking_id, amount, fee, payout_status, created_at, updated_at
`

type InsertTransactionParams struct {
	RenterID     uuid.UUID      `json:"renter_id"`
	OwnerID      uuid.UUID      `json:"owner_id"`
	BookingID    uuid.UUID      `json:"booking_id"`
	Amount       pgtype.Numeric `json:"amount"`
	Fee          pgtype.Numeric `json:"fee"`
	PayoutStatus PayoutStatus   `json:"payout_status"`
}

func (q *Queries) InsertTransaction(ctx context.Context, arg InsertTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, insertTransaction,
		arg.RenterID,
		arg.OwnerID,
		arg.BookingID,
		arg.Amount,
		arg.Fee,
		arg.PayoutStatus,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.RenterID,
		&i.OwnerID,
		&i.BookingID,
		&i.Amount,
		&i.Fee,
		&i.PayoutStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTransactions = `-- name: ListTransactions :many
SELECT id, renter_id, owner_id, booking_id, amount, fee, payout_status FROM transactions ORDER BY created_at
`

type ListTransactionsRow struct {
	ID           uuid.UUID      `json:"id"`
	RenterID     uuid.UUID      `json:"renter_id"`
	OwnerID      uuid.UUID      `json:"owner_id"`
	BookingID    uuid.UUID      `json:"booking_id"`
	Amount       pgtype.Numeric `json:"amount"`
	Fee          pgtype.Numeric `json:"fee"`
	PayoutStatus PayoutStatus   `json:"payout_status"`
}

func (q *Queries) ListTransactions(ctx context.Context) ([]ListTransactionsRow, error) {
	rows, err := q.db.Query(ctx, listTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTransactionsRow
	for rows.Next() {
		var i ListTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.RenterID,
			&i.OwnerID,
			&i.BookingID,
			&i.Amount,
			&i.Fee,
			&i.PayoutStatus,
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

const updateTransaction = `-- name: UpdateTransaction :one
UPDATE transactions SET payout_status = $2 WHERE id = $1 RETURNING id, renter_id, owner_id, booking_id, amount, fee, payout_status, created_at, updated_at
`

type UpdateTransactionParams struct {
	ID           uuid.UUID    `json:"id"`
	PayoutStatus PayoutStatus `json:"payout_status"`
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, updateTransaction, arg.ID, arg.PayoutStatus)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.RenterID,
		&i.OwnerID,
		&i.BookingID,
		&i.Amount,
		&i.Fee,
		&i.PayoutStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
