package unit

import (
	"context"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/entity"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAuthServiceLogin(t *testing.T) {
	db, _, errDBMock := sqlmock.New()
	assert.NoError(t, errDBMock)

	defer db.Close()

	userRepositoryMock := new(UserRepositoryMock)
	authService := service.NewAuthService(db, userRepositoryMock, validatorMock)

	loginRequest := request.UserLoginRequest{
		Username: "apollo",
		Password: "secret",
	}

	hashedPassword, _ := helper.HashPassword("secret")

	ctx := context.Background()
	validatorMock.On("StructCtx", ctx, loginRequest).Return(nil)

	expectedUser := entity.User{
		Id:          1,
		Username:    "apollo",
		Password:    hashedPassword,
		Name:        "Apollo",
		Email:       "apollo@example.xyz",
		PhoneNumber: "081746219124",
		CreatedAt:   "2020-10-10 10:10:10",
		UpdatedAt:   "2020-10-10 10:10:10",
	}
	userRepositoryMock.On("GetByUsername", ctx, db, loginRequest.Username).Return(expectedUser, nil)

	loginResponse, errLogin := authService.Login(ctx, loginRequest)
	assert.NoError(t, errLogin)

	assert.Equal(t, expectedUser.Id, loginResponse.Id)
	assert.Equal(t, expectedUser.Username, loginResponse.Username)
	assert.Equal(t, expectedUser.Name, loginResponse.Name)
	assert.Equal(t, expectedUser.Email, loginResponse.Email)
	assert.Equal(t, expectedUser.PhoneNumber, loginResponse.PhoneNumber)
	assert.Equal(t, expectedUser.CreatedAt, loginResponse.CreatedAt)
	assert.NotEmpty(t, loginResponse.AccessToken)
	assert.NotEmpty(t, loginResponse.AccessToken)
}
