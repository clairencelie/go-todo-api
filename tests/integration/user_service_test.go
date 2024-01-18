package integration

import (
	"context"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/service"
	testhelper "go_todo_api/tests/test_helper"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceInitialize(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())

	assert.NotNil(t, userService)
}

func TestUserServiceCreate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userCreateRequest := request.UserCreateRequest{
		Username:    "budi",
		Password:    "rahasia",
		Name:        "Budi",
		Email:       "budi@example.xyz",
		PhoneNumber: "08767890123",
	}

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())

	err := userService.Create(context.Background(), userCreateRequest)

	assert.Nil(t, err)
}

func TestUserServiceFindById(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userLastInsertId := testhelper.InsertSingleUser(db)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())

	user, err := userService.Find(context.Background(), int(userLastInsertId))

	assert.Nil(t, err)

	assert.NotNil(t, user)
}

func TestUserServiceUpdate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userLastInsertId := testhelper.InsertSingleUser(db)

	userUpdateRequest := request.UserUpdateRequest{
		Id:          int(userLastInsertId),
		Username:    "budiman",
		Name:        "Budiman",
		Email:       "budiman@example.xyz",
		PhoneNumber: "0512345",
	}

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())

	err := userService.Update(context.Background(), userUpdateRequest)

	assert.Nil(t, err)
}

func TestUserServiceRemove(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userLastInsertId := testhelper.InsertSingleUser(db)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())

	err := userService.Remove(context.Background(), int(userLastInsertId))

	assert.Nil(t, err)
}
