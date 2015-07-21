package db

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(connectionString string) (err error) {
	DB, err = sqlx.Connect("postgres", connectionString)
	return
}
