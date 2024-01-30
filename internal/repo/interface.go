package repo

import (
	"github.com/jmoiron/sqlx"
	"infotecs/internal/model"
)

type PostgresDb struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) WalletRepo {
	return &PostgresDb{
		db: db,
	}
}

type WalletRepo interface {
	Wallet
}

type Wallet interface {
	Create(id string) (*model.Wallet, error)
	ChangeBalance(parameters *model.ParametersTransaction) error
	GetHistoryWallet(id string) ([]*model.ParametersTransaction, error)
	GetWalletInfo(idWallet string) (*model.Wallet, error)
	CheckValidId(id string) error
}
