package notification

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	SEL_NOTIFICATIONS_SINCE   = `SELECT * from notifications WHERE created_on >= $1 AND heard = $2`
	SEL_NOTIFICATIONS_UNHEARD = `SELECT * from notifications WHERE heard = %d`
)

func GetSince(conn *sqlx.DB, since time.Time) (n []*Notification, err error) {
	n = make([]*Notification, 0)
	_since := since.Format("2006-01-02 15:04:05")
	err = conn.Select(&n, SEL_NOTIFICATIONS_SINCE, _since, HRD_UNHEARD)
	return
}

func GetUnheard(conn *sqlx.DB) (n []*Notification, err error) {
	n = make([]*Notification, 0)
	err = conn.Select(&n, fmt.Sprintf(SEL_NOTIFICATIONS_UNHEARD, HRD_UNHEARD))
	return
}
