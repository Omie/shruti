package setting

import (
	"github.com/jmoiron/sqlx"
)

const (
	UPDATE_SETTING = `UPDATE settings
	SET
	value=:value
	WHERE
	key=:key
	AND
	provider=(SELECT id from providers WHERE name=:name)`
)

func (self *Setting) Update(conn *sqlx.DB) (err error) {
	err = self.ValidUpdate()
	if err != nil {
		return
	}
	rows, err := conn.NamedQuery(UPDATE_SETTING, self)
	if err == nil {
		defer rows.Close()
	}

	return
}
