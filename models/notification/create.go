package notification

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	INSERT_NOTIFICATION = `INSERT INTO notifications 
	(title, url, key, provider, priority, action, created_on)
	VALUES 
	(:title, :url, :key, (SELECT id FROM providers WHERE name=:provider), :priority, :action, :created_on)`
)

func (self *Notification) Insert(conn *sqlx.DB) (err error) {
	err = self.ValidInsert()
	if err != nil {
		return
	}
	self.CreatedOn = time.Now().In(time.UTC)
	_, err = conn.NamedQuery(INSERT_NOTIFICATION, self)
	return
}
