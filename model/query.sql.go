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

const getOwner = `-- name: GetOwner :one
SELECT id, first_name, middle_name, last_name, email, phone_number, account_number, bank_name FROM owner
WHERE id = $1 LIMIT 1
`

type GetOwnerRow struct {
	ID            uuid.UUID   `json:"id"`
	FirstName     string      `json:"firstName"`
	MiddleName    pgtype.Text `json:"middleName"`
	LastName      string      `json:"lastName"`
	Email         string      `json:"email"`
	PhoneNumber   string      `json:"phoneNumber"`
	AccountNumber string      `json:"accountNumber"`
	BankName      string      `json:"bankName"`
}

func (q *Queries) GetOwner(ctx context.Context, id uuid.UUID) (GetOwnerRow, error) {
	row := q.db.QueryRow(ctx, getOwner, id)
	var i GetOwnerRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.AccountNumber,
		&i.BankName,
	)
	return i, err
}

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
