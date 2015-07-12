package notification

import (
	//"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	SEL_NOTIFICATIONS_SINCE = `SELECT * from notifications WHERE created_on >= $1`
)

func GetSince(conn *sqlx.DB, since time.Time) (n []*Notification, err error) {
	n = make([]*Notification, 0)
	_since := since.Format("2006-01-02 15:04:05")
	err = conn.Select(&n, SEL_NOTIFICATIONS_SINCE, _since)
	return
}
