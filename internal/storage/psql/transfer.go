package psql

import (
	"bina/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransferPostgres struct {
	db *sqlx.DB

}

func (t TransferPostgres) CreateTransfer(transfer *core.Transfer) (int, error) {
//	VIA TRANSACTION


//	::: transaction setup

	tx, err := t.db.Begin()
	if err != nil {
		return -1, fmt.Errorf("while tx trans :::%w",err)
	}


//	::: TRANSFER  create

	var id int
	q:= fmt.Sprintf(`insert into %s (from_account_id, to_account_id, amount, currency ) values ($1,$2,$3,$4) returning id`,TransferTable)
	err = tx.QueryRow(q, transfer.FromAccountID, transfer.ToAccountID, transfer.Amount, transfer.Currency).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("while insert :::%w",err)
	}
//	::: ACCOUNT FROM  update

	q= fmt.Sprintf(`update  %s set balance = balance-$1 where id = $2 `,AccountsTable)
	_,err = tx.Exec(q, transfer.Amount,  transfer.FromAccountID)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return -1, fmt.Errorf("while rollback :::%w" ,err2)
		}
		return -1, fmt.Errorf("while updating from acc :::%w" ,err)
	}

	if err != nil {
		return -1, fmt.Errorf("while insert :::%w",err)
	}
//	::: ACCOUNT TO  update

	q= fmt.Sprintf(`update  %s set balance = balance+$1 where id = $2 `,AccountsTable)
	_,err = tx.Exec(q, transfer.Amount,  transfer.ToAccountID)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return -1, fmt.Errorf("while rollback :::%w" ,err2)
		}
		return -1, fmt.Errorf("while updating to acc  :::%w" ,err)
	}

//	::: commit setup
	err = tx.Commit()
	if err != nil {
		return -1, err
	}
	return id,nil
}

	func (t TransferPostgres) GetTransferById(transferId int) (*core.Transfer, error) {
		q:=fmt.Sprintf(`select id, from_account_id, to_account_id, amount, currency from %s where id = $1`,TransferTable)
		var tra = new(core.Transfer)
		err := t.db.QueryRow(q, transferId).Scan(&tra.Id, &tra.FromAccountID, &tra.ToAccountID, &tra.Amount, &tra.Currency, )
		if err != nil {
			return nil, err
		}

		return tra,err

	}

func (t TransferPostgres) GetTransfers(userID int) ([]*core.Transfer, error) {
	q:=fmt.Sprintf(`select id, from_account_id, to_account_id, amount, currency from %s where from_account_id = $1`,TransferTable)
	var transfers =make([]*core.Transfer,0)
	err := t.db.Select(&transfers,q, userID)
	if err != nil {
		return nil, err
	}

	return transfers,err
}

func NewTransferPostgres(db *sqlx.DB) *TransferPostgres {
return &TransferPostgres{db}
}
