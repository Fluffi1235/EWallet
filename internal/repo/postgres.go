package repo

import (
	"github.com/jmoiron/sqlx"
	"infotecs/internal/model"
)

const (
	walletTable       = "wallets"
	transactionsTable = "transactions_history"
)

var TimeoutTx int

func NewPostgresDB(cfg *model.Config) (*sqlx.DB, error) {
	TimeoutTx = cfg.TimeoutTx
	db, err := sqlx.Connect("postgres", cfg.DataDB)
	if err != nil {
		return nil, err
	}
	return db, nil
}
