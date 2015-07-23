package notification

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	INSERT_NOTIFICATION = `INSERT INTO notifications 
	(title, url, key, provider, priority, action, created_on)
	VALUES 
	(:title, :url, :key, (SELECT id FROM providers WHERE name=:provider_name), :priority, :action, :created_on)`
)

func (self *Notification) Insert(conn *sqlx.DB) (err error) {
	err = self.ValidInsert()
	if err != nil {
		return
	}
	self.CreatedOn = time.Now().In(time.UTC)
	rows, err := conn.NamedQuery(INSERT_NOTIFICATION, self)
	if err == nil {
		defer rows.Close()
	}

	return
}
