package psql

import (
	"bina/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)


func TestTransferPostgres_CreateTransfer(t1 *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		transfer *core.Transfer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "createONE",
			fields:  fields{
				db: db,
			},
			args:    args{
				transfer: &core.Transfer{
					Id:            1,
					FromAccountID: 1,
					ToAccountID:   2,
					Amount:        800,
					Currency:      "Dollar",
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TransferPostgres{
				db: tt.fields.db,
			}
			got, err := t.CreateTransfer(tt.args.transfer)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("Transfer successed!!! --->>>>> ",got)
		})
	}
}

func TestTransferPostgres_GetTransferById(t1 *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		transferId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.Transfer
		wantErr bool
	}{

		{
			name:    "getONE",
			fields:  fields{
				db: db,
			},
			args:    args{
				transferId: 2,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TransferPostgres{
				db: tt.fields.db,
			}
			got, err := t.GetTransferById(tt.args.transferId)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTransferById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("\nTRANSFER --->>>> ",got)
		})
	}
}

func TestTransferPostgres_GetTransfers(t1 *testing.T) {
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
		want    []*core.Transfer
		wantErr bool
	}{
		{
			name:    "GETall",
			fields:  fields{
				db: db,
			},
			args:    args{
				userID: 1,
			},
		},

	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TransferPostgres{
				db: tt.fields.db,
			}
			got, err := t.GetTransfers(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTransfers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, transfer := range got {
				fmt.Println("\nAnother one Transer :::: ",transfer)

			}
		})
	}
}
