package helper

import (
	"database/sql"
)

func CheckRowsAffected(sqlResult sql.Result) error {
	rowsAffected, err := sqlResult.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRowsNotAffected
	}

	return nil
}
