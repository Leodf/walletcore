-- name: FindBalancesByAccountId :many
SELECT id, account_id, amount, last_update FROM balances WHERE account_id = ?;

-- name: SaveBalances :exec
INSERT INTO balances (id, account_id, amount, last_update) VALUES
(?, ?, ?, ?),
(?, ?, ?, ?);

-- name: CheckAccountIdExists :one
SELECT EXISTS( SELECT id, account_id, amount, last_update FROM balances WHERE account_id = ?);