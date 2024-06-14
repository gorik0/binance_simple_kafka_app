package service

import (
	"bina/internal/core"
	"bina/internal/es"
	"bina/internal/storage"
	"context"
)

type TransferService struct {
	TransferRepo  storage.Transfer
	AccountRepo   storage.Account
	messageBroker es.TransferMessageBroker
}

func (t TransferService) CreateTransfer(transfer *core.Transfer) (int, error) {
	if err := t.validTransfer(transfer); err != nil {
		return -1, err
	}
	id, err := t.TransferRepo.CreateTransfer(transfer)
	if err != nil {
		return -1, err
	}
	err = t.messageBroker.PublishTransferCreate(context.Background(), *transfer)
	if err != nil {
		return -1, err
	}
	return id,nil
}

func (t TransferService) GetTransferById(transferId int) (*core.Transfer, error) {
	transfer, err := t.TransferRepo.GetTransferById(transferId)
	if err != nil {
		return nil, err
	}
	return transfer, nil
}

func (t TransferService) GetTransfers(userID int) ([]*core.Transfer, error) {
	transfer, err := t.TransferRepo.GetTransfers(userID)
	if err != nil {
		return nil, err
	}
	return transfer, nil
}

func (t TransferService) validTransfer(transfer *core.Transfer) error {

	if transfer.Amount <= 0 {
		return core.ErrAmountMustBePositive
	}
	if transfer.ToAccountID == transfer.ToAccountID {
		return core.ErrSameAccount
	}

	if _, err := t.validAccount(transfer.FromAccountID); err != nil {
		return err
	}
	if _, err := t.validAccount(transfer.ToAccountID); err != nil {
		return err
	}
	return nil
}

func (t TransferService) validAccount(id int) (*core.Account, error) {
	acc, err := t.AccountRepo.GetAccountById(id)
	if err != nil {
		return nil, err
	}
	return acc, nil

}

func NewTransferService(transfRepo storage.Transfer, accRepo storage.Account, msgBr es.MessageBroker) *TransferService {
	return &TransferService{
		TransferRepo: transfRepo,
		AccountRepo:  accRepo,
	}

}
