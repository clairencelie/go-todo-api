package integration

import (
	"context"
	"database/sql"
	"go_todo_api/database"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	testhelper "go_todo_api/tests/test_helper"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
)

func setupDb() (*sql.DB, error) {
	db, errDbConn := database.NewDB("./../../", true)

	if errDbConn != nil {
		return nil, errDbConn
	}

	testhelper.ResetDB(db)

	return db, nil
}

func TestUserRepositoryInitialize(t *testing.T) {
	userRepository := repository.NewUserRepository()

	assert.NotNil(t, userRepository)
}

func TestUserRepositoryGetById(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userLastInsertId := testhelper.InsertSingleUser(db)

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	user, err := userRepository.Get(ctx, db, int(userLastInsertId))

	assert.Nil(t, err)

	assert.Equal(t, int(userLastInsertId), user.Id)
	assert.Equal(t, "budi", user.Username)
	assert.Equal(t, "Budi", user.Name)
	assert.Equal(t, "budi@example.xyz", user.Email)
	assert.Equal(t, "081234567", user.PhoneNumber)
}

func TestUserRepositoryInsert(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	userCreateRequest := &request.UserCreateRequest{
		Username:    "budi",
		Password:    "rahasia",
		Name:        "Budi",
		Email:       "budi@example.xyz",
		PhoneNumber: "087654321",
	}

	errInsert := userRepository.Insert(ctx, db, *userCreateRequest)
	assert.Nil(t, errInsert)
}

func TestUserRepositoryUpdate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userLastInsertId := testhelper.InsertSingleUser(db)

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	userUpdateRequest := request.UserUpdateRequest{
		Id:          int(userLastInsertId),
		Username:    "cipi",
		Name:        "Chipi Chipi Chapa Chapa",
		Email:       "lubilubi@example.xyz",
		PhoneNumber: "0781724829",
	}

	err := userRepository.Update(ctx, db, userUpdateRequest)

	assert.Nil(t, err)
}

func TestUserRepositoryDelete(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userLastInsertId := testhelper.InsertSingleUser(db)

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	tx, errTxBegin := db.Begin()

	assert.Nil(t, errTxBegin)

	err := userRepository.Delete(ctx, tx, int(userLastInsertId))

	if err == nil {
		tx.Commit()
	}

	assert.Nil(t, err)
}
