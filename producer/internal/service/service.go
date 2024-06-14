package service

import (
	"bina/internal/core"
	"bina/internal/es"
	ser "bina/internal/service/webapi"
	"bina/internal/storage"
)

type Account interface {
	CreateAccount(user *core.Account) (int, error)
	GetAccountById(accId int)(*core.Account,error)
	GetAccounts(userID int)([]*core.Account,error)
	UpdateAccount(user *core.Account) error

}
type Authorization interface {
	CreateUser(user *core.User)(int,error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string)(int,error)

}
type Transfer interface {
	CreateTransfer(transfer *core.Transfer)(int,error)
	GetTransferById(transferId int)(*core.Transfer,error)
	GetTransfers(userID int)([]*core.Transfer,error)
}

type Coin interface {
	GetCoinPrice(symbol string) ([]core.SymbolPrice,error)
	
}
type Service struct {

	
	Account
	Authorization
	Transfer
	Coin
}


func NewService(storage *storage.Storage,bapi *ser.BinanceWEBapi,msgBroker *es.MessageBroker)*Service{
	return &Service{
		Account:       NewAccountService(storage.Account,*msgBroker),
		Authorization: NewAuthService(storage),
		Transfer:      NewTransferService(storage.Transfer,storage.Account,*msgBroker,),
		Coin:          NewCoinService(bapi),
	}
}