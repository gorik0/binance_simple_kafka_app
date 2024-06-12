package psql

import (
	"bina/internal/core"
	"github.com/jmoiron/sqlx"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
	
}

func (a AuthorizationPostgres) CreateUser(user *core.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorizationPostgres) GetUSer(login string, password string) *core.User {
	//TODO implement me
	panic("implement me")
}

func NewAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
return &AuthorizationPostgres{db: db}
}
