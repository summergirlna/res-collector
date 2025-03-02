package rdb

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"res-collector/internal/repository"
)

type Transaction struct {
	client *sqlx.DB
}

func NewTransaction(client *sqlx.DB) repository.Transaction {
	return &Transaction{client: client}
}

var txKey = struct{}{}

func (t Transaction) Do(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error) {
	slog.Info("start transaction begin")
	tx, err := t.client.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, &txKey, tx)

	slog.Info("start sql")
	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	slog.Info("start commit")
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return v, nil
}

func GetTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sqlx.Tx)
	return tx, ok
}
