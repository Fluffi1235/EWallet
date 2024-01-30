package repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"infotecs/internal/logger"
	"infotecs/internal/model"
	"time"
)

func (r *PostgresDb) Create(id string) (*model.Wallet, error) {
	wallet := &model.Wallet{}
	query := fmt.Sprintf("INSERT INTO %s(id) VALUES($1) RETURNING id, balance", walletTable)

	err := r.db.Get(wallet, query, id)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *PostgresDb) GetWalletInfo(idWallet string) (*model.Wallet, error) {
	wallet := &model.Wallet{}
	query := fmt.Sprintf("Select id, balance from %s where id = $1", walletTable)

	err := r.db.Get(wallet, query, idWallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *PostgresDb) ChangeBalance(parameters *model.ParametersTransaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(TimeoutTx)*time.Second)
	defer cancel()

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer commitOrRollbackTx(tx, err)

	err = sqlUpdateBalance(tx, parameters.FromId, -parameters.Amount)
	if err != nil {
		return err
	}

	err = sqlUpdateBalance(tx, parameters.ToId, parameters.Amount)
	if err != nil {
		return err
	}

	queryLogTransaction := fmt.Sprintf("insert into %s (fromId, toId, amount, time) values($1, $2, $3, $4)", transactionsTable)
	_, err = tx.Exec(queryLogTransaction, parameters.FromId, parameters.ToId, parameters.Amount, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDb) GetHistoryWallet(id string) ([]*model.ParametersTransaction, error) {
	var historyWallet []*model.ParametersTransaction
	query := fmt.Sprintf("Select fromId, toId, amount, time from %s where fromId = $1 or toId = $1", transactionsTable)

	err := r.db.Select(&historyWallet, query, id)
	if err != nil {
		return nil, err
	}

	return historyWallet, nil
}

func sqlUpdateBalance(tx *sqlx.Tx, id string, amount int) error {
	query := fmt.Sprintf("update %s set balance = balance + $1 where id = $2", walletTable)
	_, err := tx.Exec(query, amount, id)
	if err != nil {
		return err
	}

	return nil
}

func commitOrRollbackTx(tx *sqlx.Tx, err error) {
	if err != nil {
		tx.Rollback()
	} else {
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			logger.Logger.Error("Error commit tx: ", err.Error())
		}
	}
}

func (r *PostgresDb) CheckValidId(id string) error {
	var walletId string
	query := fmt.Sprintf("Select id from %s where id = $1", walletTable)

	err := r.db.Get(&walletId, query, id)
	if err != nil {
		return err
	}

	return nil
}
