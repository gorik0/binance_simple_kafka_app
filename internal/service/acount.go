package service

import (
	"bina/internal/core"
	"bina/internal/es"
	"bina/internal/storage"
	"bina/utils"
	"context"
)

type AccountService struct {
	AccountRepo storage.Account
	messageBroker es.AccountMessageBroker
}

func (a AccountService) CreateAccount(account *core.Account) (int, error) {

	if !utils.IsSupportedCoin(account.Currency) {
		return -1,core.ErrUnsupportedCurrency
	}
	id, err := a.AccountRepo.CreateAccount(account)
	err = a.messageBroker.PublishAccountCreate(context.Background(), *account)
	if err != nil {
		return -1, err
	}
	if err != nil {
		return -1, err
	}
	return id,nil



}

func (a AccountService) GetAccountById(accId int) (*core.Account, error) {
	account, err := a.AccountRepo.GetAccountById(accId)
	if err != nil {
		return nil, err
	}
	return account,nil
}

func (a AccountService) GetAccounts(userID int) ([]*core.Account, error) {
	accounts, err := a.AccountRepo.GetAccounts(userID)
	if err != nil {
		return nil, err
	}
	return accounts,nil
}

func (a AccountService) UpdateAccount(account *core.Account) error {



		err := a.AccountRepo.UpdateAccount(account)
	err = a.messageBroker.PublishAccountCreate(context.Background(), *account)
	if err != nil {
		return err
	}
		if err != nil {
			return  err
		}
		return nil



	}

func NewAccountService(accRepo storage.Account,msgBr es.MessageBroker) *AccountService {
	return &AccountService{
		AccountRepo:   accRepo,
		messageBroker: msgBr.Account,
	}
}