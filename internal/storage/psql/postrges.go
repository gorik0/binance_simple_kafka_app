package psql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
type Config struct {
	Host string
	Password string
	Port string
	DB string
	User string
	SSLMode string

}

var (
UserTable = "users"
AccountsTable = "accounts"
TransferTable = "transfers"
)



func NewPostgresDB(cfg *Config)(*sqlx.DB,error){
	connStr := fmt.Sprintf("host=%s password=%s port=%s dbname=%s sslmode=%s",cfg.Host,cfg.Password,cfg.Port,cfg.DB,cfg.SSLMode)
	db,err := sqlx.Connect("postgres",connStr)

	if err != nil {
		return nil, fmt.Errorf("-- connect  -- ::: %w",err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("-- ping  -- ::: %w",err)
	}
	return db, nil
}