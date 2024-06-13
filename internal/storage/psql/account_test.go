package psql

import (
	"bina/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestAccountPostgres_CreateAccount(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		account *core.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "createUser",
			fields: fields{
				db: db,
			},
			args: args{
				account: &core.Account{
					Id:       0,
					UserId:   1,
					Balance:  1000,
					Currency: "dolla",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountPostgres{
				db: tt.fields.db,
			}
			id, err := a.CreateAccount(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("\nGotta id :::: -> ", id)

		})
	}
}

func TestAccountPostgres_GetAccountById(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		accId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.Account
		wantErr bool
	}{
		{
			name: "geteexisted",
			fields: fields{
				db: db,
			},
			args: args{
				accId: 2,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountPostgres{
				db: tt.fields.db,
			}
			got, err := a.GetAccountById(tt.args.accId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("USER --->>> ::: ", got)
		})
	}
}

func TestAccountPostgres_GetAccounts(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*core.Account
		wantErr bool
	}{
		{
			name: "accs",
			fields: fields{
				db: db,
			},
			args: args{
				userID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountPostgres{
				db: tt.fields.db,
			}
			got, err := a.GetAccounts(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, account := range got {

				fmt.Println("\n users ---->>>> ", account)
			}
		})
	}

}

func TestAccountPostgres_UpdateAccount(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		account *core.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "updateONE",
			fields:  fields{
				db: db,
			},
			args:    args{
				account: &core.Account{
					Id:       1,
					UserId:   1,
					Balance:  999,
					Currency: "euro",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountPostgres{
				db: tt.fields.db,
			}
			if err := a.UpdateAccount(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

