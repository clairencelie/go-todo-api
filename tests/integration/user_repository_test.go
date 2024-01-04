package integration

import (
	"context"
	"go_todo_api/database"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitiateUserRepository(t *testing.T) {
	userRepository := repository.NewUserRepository()

	assert.NotNil(t, userRepository)
}

func TestGetUserById(t *testing.T) {
	db, _ := database.NewDB("./../../")

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	user, err := userRepository.Get(ctx, db, 1)

	assert.Nil(t, err)

	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "budi", user.Username)
	assert.Equal(t, "rahasia", user.Password)
	assert.Equal(t, "budi@example.xyz", user.Email)
	assert.Equal(t, "Budi Gemilang", user.Name)
}

func TestGetAllUser(t *testing.T) {
	db, _ := database.NewDB("./../../")

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	users, err := userRepository.GetAll(ctx, db)

	assert.Nil(t, err)

	assert.True(t, len(users) > 0)
}

func TestInsertUser(t *testing.T) {
	db, _ := database.NewDB("./../../")

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	tx, errTxBegin := db.Begin()

	assert.Nil(t, errTxBegin)

	userCreateRequest := &request.UserCreateRequest{
		Username:    "popon",
		Password:    "rahasia",
		Name:        "Popon Pipan",
		Email:       "popon@example.xyz",
		PhoneNumber: "0891829410",
	}

	errInsert := userRepository.Insert(ctx, tx, *userCreateRequest)
	assert.Nil(t, errInsert)
	tx.Commit()
}

func TestDeleteUser(t *testing.T) {
	db, _ := database.NewDB("./../../")

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	tx, errTxBegin := db.Begin()

	assert.Nil(t, errTxBegin)

	err := userRepository.Delete(ctx, tx, 3)

	if err == nil {
		tx.Commit()
	}

	assert.Nil(t, err)
}

func TestUpdateUser(t *testing.T) {
	db, _ := database.NewDB("./../../")

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	tx, errTxBegin := db.Begin()

	assert.Nil(t, errTxBegin)

	userUpdateRequest := request.UserUpdateRequest{
		Id:          5,
		Username:    "pupu",
		Name:        "Pupu Pipi Popo",
		Email:       "pupu@example.xyz",
		PhoneNumber: "0781724829",
	}

	err := userRepository.Update(ctx, tx, userUpdateRequest)

	if err == nil {
		tx.Commit()
	}

	assert.Nil(t, err)
}
