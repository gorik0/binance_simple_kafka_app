package psql

import (
	"bina/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

var err error

func init() {
	cfg := &Config{
		Host:     "localhost",
		Password: "1",
		Port:     "5432",
		DB:       "test",
		User:     "postgres",
		SSLMode:  "disable",
	}


	db, err = NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}
}
func TestSome(t *testing.T){
	println("After")
}

var db *sqlx.DB

func TestAuthorizationPostgres_GetUSer(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.User
		wantErr bool
	}{
		{
			name:    "get exists user",
			fields:  fields{db:db},
			args:    args{
				login: "egor",
				password: "1",

			},


		},
		{
			name:    "get not existed user",
			fields:  fields{db:db},
			args:    args{
				login: "notegor",
				password: "1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AuthorizationPostgres{
				db: tt.fields.db,
			}
			got, err := a.GetUSer(tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unexpected errror  = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("pass ::: ",got)
		})
	}
}


