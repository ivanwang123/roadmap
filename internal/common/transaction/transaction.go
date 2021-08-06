package transaction

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

const TransactionCtxKey = "transaction"

type TxFunc func(context.Context) error

type Conn interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func GetConn(ctx context.Context, db Conn) Conn {
	tx, ok := ctx.Value(TransactionCtxKey).(*sqlx.Tx)
	if ok {
		return tx
	}
	return db
}

func NewTransaction(ctx context.Context, db *sqlx.DB, fn TxFunc) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	txCtx := context.WithValue(ctx, TransactionCtxKey, tx)

	err = fn(txCtx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
