package setting

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

const (
	//SELECT_SETTING      = `SELECT * from settings WHERE key=$1`
	SELECT_SETTING      = `SELECT settings.*, providers.name from settings, providers where settings.provider = providers.id AND key=$1`
	SELECT_ALL_SETTINGS = `SELECT settings.*, providers.name from settings, providers where settings.provider = providers.id`
)

func Get(conn *sqlx.DB, key string) (s *Setting, err error) {
	if key == "" {
		return nil, errors.New("Invalid key")
	}
	s = new(Setting)
	err = conn.Get(s, SELECT_SETTING, key)
	return
}

func GetAll(conn *sqlx.DB) (s []*Setting, err error) {
	s = make([]*Setting, 0)
	err = conn.Select(&s, SELECT_ALL_SETTINGS)
	return
}
