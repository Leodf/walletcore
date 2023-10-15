// versions:
//   sqlc v1.22.0
// source: balances.sql

package database

import (
	"context"
)

const checkAccountIdExists = `-- name: CheckAccountIdExists :one
SELECT EXISTS( SELECT id, account_id, amount, last_update FROM balances WHERE account_id = ?)
`

func (q *Queries) CheckAccountIdExists(ctx context.Context, accountID string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkAccountIdExists, accountID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const findBalancesByAccountId = `-- name: FindBalancesByAccountId :many
SELECT id, account_id, amount, last_update FROM balances WHERE account_id = ?
`

func (q *Queries) FindBalancesByAccountId(ctx context.Context, accountID string) ([]Balance, error) {
	rows, err := q.db.QueryContext(ctx, findBalancesByAccountId, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Balance
	for rows.Next() {
		var i Balance
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.LastUpdate,
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

const saveBalances = `-- name: SaveBalances :exec
INSERT INTO balances (id, account_id, amount, last_update) VALUES
(?, ?, ?, ?),
(?, ?, ?, ?)
`

func (q *Queries) SaveBalances(ctx context.Context, arg SaveBalancesParams) error {
	_, err := q.db.ExecContext(ctx, saveBalances,
		arg.ID,
		arg.AccountID,
		arg.Amount,
		arg.LastUpdate,
		arg.ID_2,
		arg.AccountID_2,
		arg.Amount_2,
		arg.LastUpdate_2,
	)
	return err
}
