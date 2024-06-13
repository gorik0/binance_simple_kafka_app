package storage

import (
	"bina/internal/core"
	"bina/internal/storage/psql"
	"github.com/jmoiron/sqlx"
)

type Account interface {
	CreateAccount(user *core.Account) (int, error)
	GetAccountById(accId int)(*core.Account,error)
	GetAccounts(userID int)([]*core.Account,error)
	UpdateAccount(user *core.Account) error

}
type Authorization interface {
	CreateUser(user *core.User)(int,error)
	GetUSer(login string, password string) (*core.User, error)

}
type Transfer interface {
	CreateTransfer(transfer *core.Transfer)(int,error)
	GetTransferById(transferId int)(*core.Transfer,error)
	GetTransfers(userID int)([]*core.Transfer,error)
}
type Storage struct {
	Account
	Authorization
	Transfer
}


func NewStorage(db *sqlx.DB)*Storage{
	return &Storage{
		Account:       psql.NewAccountPostgres(db),
		Authorization: psql.NewAuthorizationPostgres(db),
		Transfer:      psql.NewTransferPostgres(db),
	}
}

