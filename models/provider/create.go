package provider

import (
	"github.com/jmoiron/sqlx"
)

const (
	INSERT_PROVIDER = `INSERT INTO providers 
	(name, display_name, description, web_url, icon_url, voice)
	VALUES 
	(:name, :display_name, :description, :web_url, :icon_url, :voice)`
)

func (self *Provider) Insert(conn *sqlx.DB) (err error) {
	err = self.ValidInsert()
	if err != nil {
		return
	}
	rows, err := conn.NamedQuery(INSERT_PROVIDER, self)
	if err == nil {
		defer rows.Close()
	}

	return
}
