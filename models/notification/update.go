package notification

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	MARK_HEARD_NOTIFICATION = `UPDATE notifications
	SET
	heard=%d
	WHERE
	id in %s`
)

func MarkHeard(conn *sqlx.DB, unheardIds []int) (err error) {
	if len(unheardIds) < 1 {
		return errors.New("Empty list of ids")
	}

	ids, err := json.Marshal(unheardIds)
	if err != nil {
		return
	}

	_strIds := fmt.Sprintf("%s", ids)
	strIds := "(" + _strIds[1:len(_strIds)-1] + ")"

	_, err = conn.NamedExec(fmt.Sprintf(MARK_HEARD_NOTIFICATION, HRD_UNHEARD, strIds), new(struct{}))
	if err != nil {
		return
	}

	return
}
