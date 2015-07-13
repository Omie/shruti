package notification

import (
	"github.com/jmoiron/sqlx"
)

const (
	INSERT_NOTIFICATION = `INSERT INTO notifications 
	(title, url, key, provider, priority, action)
	VALUES 
	(:title, :url, :key, (SELECT id FROM providers WHERE name=:provider), :priority, :action)`
)

func (self *Notification) Insert(conn *sqlx.DB) (err error) {
	err = self.ValidInsert()
	if err != nil {
		return
	}
	_, err = conn.NamedQuery(INSERT_NOTIFICATION, self)
	return
}
