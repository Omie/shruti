package provider

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

const (
	SELECT_PROVIDER = `SELECT * from providers WHERE name=$1`
)

func Get(conn *sqlx.DB, pname string) (p *Provider, err error) {
	if pname == "" {
		return nil, errors.New("Invalid provider name")
	}
	p = new(Provider)
	err = conn.Get(p, SELECT_PROVIDER, pname)
	return
}
