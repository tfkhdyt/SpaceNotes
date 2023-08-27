// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: refresh_token_query.sql

package sqlc

import (
	"context"
)

const addToken = `-- name: AddToken :exec
INSERT INTO refresh_tokens (token) VALUES ($1)
`

func (q *Queries) AddToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, addToken, token)
	return err
}

const deleteToken = `-- name: DeleteToken :exec
DELETE FROM refresh_tokens WHERE token = $1
`

func (q *Queries) DeleteToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteToken, token)
	return err
}

const findToken = `-- name: FindToken :one
SELECT token FROM refresh_tokens WHERE token = $1
`

func (q *Queries) FindToken(ctx context.Context, token string) (string, error) {
	row := q.db.QueryRow(ctx, findToken, token)
	err := row.Scan(&token)
	return token, err
}
