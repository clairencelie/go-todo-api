package helper

import (
	"database/sql"
	"errors"
)

var (
	errRowsNotAffected = errors.New("no rows affected")
)

func CheckRowsAffected(sqlResult sql.Result) error {
	rowsAffected, err := sqlResult.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errRowsNotAffected
	}

	return nil
}
