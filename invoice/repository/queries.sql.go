// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package repository

import (
	"context"

	"github.com/shopspring/decimal"
)

const createInvoice = `-- name: CreateInvoice :one
INSERT INTO invoices(user_id, service_id, order_id, amount)
VALUES ($1, $2, $3, $4)
RETURNING id, service_id, order_id, user_id, amount, payment_status, created_at, updated_at
`

type CreateInvoiceParams struct {
	UserID    int64           `json:"user_id"`
	ServiceID int64           `json:"service_id"`
	OrderID   int64           `json:"order_id"`
	Amount    decimal.Decimal `json:"amount"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, createInvoice,
		arg.UserID,
		arg.ServiceID,
		arg.OrderID,
		arg.Amount,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.ServiceID,
		&i.OrderID,
		&i.UserID,
		&i.Amount,
		&i.PaymentStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getInvoiceByID = `-- name: GetInvoiceByID :one
SELECT id, service_id, order_id, user_id, amount, payment_status, created_at, updated_at
FROM invoices
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetInvoiceByID(ctx context.Context, id int64) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, getInvoiceByID, id)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.ServiceID,
		&i.OrderID,
		&i.UserID,
		&i.Amount,
		&i.PaymentStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getInvoicesByUserID = `-- name: GetInvoicesByUserID :many
SELECT id, service_id, order_id, user_id, amount, payment_status, created_at, updated_at
FROM invoices
WHERE user_id = $1
    AND id > $2
LIMIT $3
`

type GetInvoicesByUserIDParams struct {
	UserID int64 `json:"user_id"`
	ID     int64 `json:"id"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) GetInvoicesByUserID(ctx context.Context, arg GetInvoicesByUserIDParams) ([]Invoice, error) {
	rows, err := q.db.QueryContext(ctx, getInvoicesByUserID, arg.UserID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Invoice
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.ServiceID,
			&i.OrderID,
			&i.UserID,
			&i.Amount,
			&i.PaymentStatus,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStatus = `-- name: UpdateStatus :one
UPDATE invoices
SET payment_status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, service_id, order_id, user_id, amount, payment_status, created_at, updated_at
`

type UpdateStatusParams struct {
	ID            int64       `json:"id"`
	PaymentStatus ValidStatus `json:"payment_status"`
}

func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, updateStatus, arg.ID, arg.PaymentStatus)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.ServiceID,
		&i.OrderID,
		&i.UserID,
		&i.Amount,
		&i.PaymentStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
