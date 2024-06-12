package psql

import (
	"bina/internal/core"
	"github.com/jmoiron/sqlx"
)

type TransferPostgres struct {
	db *sqlx.DB

}

func (t TransferPostgres) CreateTransfer(transfer *core.Transfer) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t TransferPostgres) GetTransferById(transferId int) (*core.Transfer, error) {
	//TODO implement me
	panic("implement me")
}

func (t TransferPostgres) GetTransfers(userID int) ([]*core.Transfer, error) {
	//TODO implement me
	panic("implement me")
}

func NewTransferPostgres(db *sqlx.DB) *TransferPostgres {
return &TransferPostgres{db}
}
