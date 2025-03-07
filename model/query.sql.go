// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const listOwner = `-- name: ListOwner :many
SELECT id, first_name, middle_name, last_name, email, phone_number,account_number, bank_name FROM owner
ORDER BY email
`

type ListOwnerRow struct {
	ID            uuid.UUID   `json:"id"`
	FirstName     string      `json:"firstName"`
	MiddleName    pgtype.Text `json:"middleName"`
	LastName      string      `json:"lastName"`
	Email         string      `json:"email"`
	PhoneNumber   string      `json:"phoneNumber"`
	AccountNumber string      `json:"accountNumber"`
	BankName      string      `json:"bankName"`
}

func (q *Queries) ListOwner(ctx context.Context) ([]ListOwnerRow, error) {
	rows, err := q.db.Query(ctx, listOwner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListOwnerRow
	for rows.Next() {
		var i ListOwnerRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Email,
			&i.PhoneNumber,
			&i.AccountNumber,
			&i.BankName,
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

const listOwnerFull = `-- name: ListOwnerFull :many
SELECT id, first_name, middle_name, last_name, email, phone_number, account_number, bank_name, created_at, updated_at FROM owner
ORDER BY email
`

func (q *Queries) ListOwnerFull(ctx context.Context) ([]Owner, error) {
	rows, err := q.db.Query(ctx, listOwnerFull)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Owner
	for rows.Next() {
		var i Owner
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Email,
			&i.PhoneNumber,
			&i.AccountNumber,
			&i.BankName,
			&i.CreatedAt,
			&i.UpdatedAt,
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
