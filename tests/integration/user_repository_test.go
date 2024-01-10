package integration

import (
	"context"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
)

func TestInitiateUserRepository(t *testing.T) {
	userRepository := repository.NewUserRepository()

	assert.NotNil(t, userRepository)
}

func TestGetUserById(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	user, err := userRepository.Get(ctx, TestDb, int(userLastInsertId))

	assert.Nil(t, err)

	assert.Equal(t, int(userLastInsertId), user.Id)
	assert.Equal(t, "budi", user.Username)
	assert.Equal(t, "rahasia", user.Password)
	assert.Equal(t, "Budi", user.Name)
	assert.Equal(t, "budi@example.xyz", user.Email)
	assert.Equal(t, "081234567", user.PhoneNumber)
}

func TestGetAllUser(t *testing.T) {
	ResetDB()

	InsertManyUser(TestDb, 5)

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	users, err := userRepository.GetAll(ctx, TestDb)

	assert.Nil(t, err)

	assert.True(t, len(users) > 0)
	assert.Len(t, users, 5)
}

func TestInsertUser(t *testing.T) {
	ResetDB()

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	tx, errTxBegin := TestDb.Begin()

	assert.Nil(t, errTxBegin)

	userCreateRequest := &request.UserCreateRequest{
		Username:    "budi",
		Password:    "rahasia",
		Name:        "Budi",
		Email:       "budi@example.xyz",
		PhoneNumber: "087654321",
	}

	errInsert := userRepository.Insert(ctx, tx, *userCreateRequest)
	assert.Nil(t, errInsert)
	tx.Commit()
}

func TestUpdateUser(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	tx, errTxBegin := TestDb.Begin()

	assert.Nil(t, errTxBegin)

	userUpdateRequest := request.UserUpdateRequest{
		Id:          int(userLastInsertId),
		Username:    "cipi",
		Name:        "Chipi Chipi Chapa Chapa",
		Email:       "lubilubi@example.xyz",
		PhoneNumber: "0781724829",
	}

	err := userRepository.Update(ctx, tx, userUpdateRequest)

	if err == nil {
		tx.Commit()
	}

	assert.Nil(t, err)
}

func TestDeleteUser(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	tx, errTxBegin := TestDb.Begin()

	assert.Nil(t, errTxBegin)

	err := userRepository.Delete(ctx, tx, int(userLastInsertId))

	if err == nil {
		tx.Commit()
	}

	assert.Nil(t, err)
}
