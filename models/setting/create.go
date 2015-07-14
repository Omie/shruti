package setting

import (
	"github.com/jmoiron/sqlx"
)

const (
	INSERT_SETTING = `INSERT INTO settings
	(key, value, provider) 
	VALUES 
	(:key, :value, (SELECT id from providers WHERE name=:name))`
)

func (self *Setting) Insert(conn *sqlx.DB) (err error) {
	err = self.ValidInsert()
	if err != nil {
		return
	}
	_, err = conn.NamedQuery(INSERT_SETTING, self)
	return
}
