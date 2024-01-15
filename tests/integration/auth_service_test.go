package integration

import (
	"context"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/service"
	testhelper "go_todo_api/tests/test_helper"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestAuthServiceInitialize(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(db, userRepository, validator.New())

	assert.NotNil(t, authService)
}

func TestAuthServiceLogin(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	testhelper.InsertSingleUser(db)

	userLoginRequest := request.UserLoginRequest{
		Username: "budi",
		Password: "rahasia",
	}

	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(db, userRepository, validator.New())

	userResponse, err := authService.Login(context.Background(), userLoginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, userResponse)
}
