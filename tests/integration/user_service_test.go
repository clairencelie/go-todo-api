package integration

import (
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeUserService(t *testing.T) {
	userService := service.NewUserService(TestDb, UserRepository, Validator)

	assert.NotNil(t, userService)
}

func TestServiceCreateUser(t *testing.T) {
	userCreateRequest := request.UserCreateRequest{
		Username:    "budi",
		Password:    "rahasia",
		Name:        "Budi",
		Email:       "budi@example.xyz",
		PhoneNumber: "08767890123",
	}

	err := UserService.Create(Ctx, userCreateRequest)

	assert.Nil(t, err)
}

func TestServiceFindUserById(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	user, err := UserService.Find(Ctx, int(userLastInsertId))

	assert.Nil(t, err)

	assert.NotNil(t, user)
}

func TestServiceFindAllUser(t *testing.T) {
	ResetDB()

	InsertManyUser(TestDb, 5)

	users, err := UserService.FindAll(Ctx)

	assert.Nil(t, err)
	assert.Greater(t, len(users), 0)
	assert.Len(t, users, 5)
}

func TestServiceUpdateUser(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	userUpdateRequest := request.UserUpdateRequest{
		Id:          int(userLastInsertId),
		Username:    "budiman",
		Name:        "Budiman",
		Email:       "budiman@example.xyz",
		PhoneNumber: "0512345",
	}

	err := UserService.Update(Ctx, userUpdateRequest)

	assert.Nil(t, err)
}

func TestServiceRemoveUser(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	err := UserService.Remove(Ctx, int(userLastInsertId))

	assert.Nil(t, err)
}
