package psql

import (
	"bina/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
	
}

func (a AuthorizationPostgres) CreateUser(user *core.User) (int, error) {
	q:=fmt.Sprintf(`INSERT INTO %s (name,username,password) values ($1, $2, $3) returning id`,UserTable)
	var id int
	err := a.db.QueryRow(q, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return -1,fmt.Errorf("creating acc ::: %w",err)
	}
	return id,nil
}

func (a AuthorizationPostgres) GetUSer(login string, password string) (*core.User, error) {


	q:=fmt.Sprintf(`select id, name,password from %s where username = $1 and password = $2`,UserTable)
	var user = new(core.User)
	err := a.db.QueryRow(q, login,password).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user,err

}

func NewAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
return &AuthorizationPostgres{db: db}
}
