package provider

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	DELETE_PROVIDER = `DELETE from providers WHERE name='%s'`
)

func Delete(conn *sqlx.DB, pname string) (err error) {
	if pname == "" || pname == "system" {
		return errors.New("Invalid provider name")
	}

	_, err = conn.NamedExec(fmt.Sprintf(DELETE_PROVIDER, pname), new(struct{}))
	return
}
