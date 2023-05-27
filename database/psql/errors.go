package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/todo/database/errors"
	"github.com/lib/pq"
)

func parseErrorType(err error) (r string) {
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation":
			r = errors.ALREADY_EXISTS
		default:
			r = errors.UNKNOWN
		}
	}
	if err == sql.ErrNoRows {
		r = errors.NOT_FOUND
	}
	return
}

func errorInRow(table, action string, err error) error {
	return errors.NewError(
		parseErrorType(err),
		fmt.Sprintf("failed to %s a row in %s table", action, table),
		err.Error(),
	)
}

func errorInRows(table, action string, err error) error {
	return errors.NewError(
		parseErrorType(err),
		fmt.Sprintf("failed to %s rows in %s table", action, table),
		err.Error(),
	)
}
