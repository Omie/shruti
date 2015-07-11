package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(user, dbname string) (err error) {
	DB, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname))
	return
}
