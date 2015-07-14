package setting

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	DELETE_SETTING = `DELETE from settings WHERE key='%s' AND provider != (SELECT id from providers WHERE name='SYSTEM')`
)

func Delete(conn *sqlx.DB, key string) (err error) {
	_, err = conn.NamedExec(fmt.Sprintf(DELETE_SETTING, key), new(struct{}))
	return
}
