package provider

import (
	"github.com/jmoiron/sqlx"
)

const (
	UPDATE_PROVIDER = `UPDATE providers
	SET
	display_name=:display_name,
	description=:description,
	web_url=:web_url,
	icon_url=:icon_url,
	active=:active,
	voice=:voice
	WHERE
	name=:name`
)

func (self *Provider) Update(conn *sqlx.DB) (err error) {
	err = self.ValidUpdate()
	if err != nil {
		return
	}
	rows, err := conn.NamedQuery(UPDATE_PROVIDER, self)
	if err == nil {
		defer rows.Close()
	}

	return
}
