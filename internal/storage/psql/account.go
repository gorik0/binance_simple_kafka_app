package psql

import (
	"bina/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func (a AccountPostgres) CreateAccount(account *core.Account) (int, error) {
	q:=fmt.Sprintf("INSERT INTO %s (user_id,balance,currency) values ($1,$2,$3) returning id",AccountsTable)
	var id int
	err := a.db.QueryRow(q, account.UserId, account.Balance, account.Currency).Scan(&id)
	if err != nil {
		return -1,fmt.Errorf("creating acc ::: %w",err)
	}
	return id,nil

}

func (a AccountPostgres) GetAccountById(accId int) (*core.Account, error) {

	q:=fmt.Sprintf(`select id, user_id, balance, currency from %s where id = $1`,AccountsTable)
	var acc = new(core.Account)
	err := a.db.QueryRow(q, accId).Scan(&acc.Id, &acc.UserId, &acc.Balance, &acc.Currency)
	if err != nil {
		return nil, err
	}

	return acc,err

}

func (a AccountPostgres) GetAccounts(userID int) ([]*core.Account, error) {
	q:=fmt.Sprintf(`select id, user_id, balance, currency from %s where user_id = $1`,AccountsTable)
	var accounts =make([]*core.Account,0)
	err := a.db.Select(&accounts,q, userID)
	if err != nil {
		return nil, err
	}

	return accounts,err
}

func (a AccountPostgres) UpdateAccount(account *core.Account) error {
	q:=fmt.Sprintf(`update %s set balance = $1 where id = $2`,AccountsTable)

	_, err := a.db.Exec(q, account.Balance,account.Id)
	return err
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres{
return &AccountPostgres{db: db}
}
