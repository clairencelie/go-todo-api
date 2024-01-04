package database

import (
	"go_todo_api/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := database.NewDB("./../../")

	assert.Nil(t, err)
	assert.NotNil(t, db)
}
