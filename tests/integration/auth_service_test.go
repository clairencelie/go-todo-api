package integration

import (
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeAuthService(t *testing.T) {
	authService := service.NewAuthService(TestDb, UserRepository, Validator)

	assert.NotNil(t, authService)
}

func TestServiceLogin(t *testing.T) {
	ResetDB()

	InsertSingleUser(TestDb)

	userLoginRequest := request.UserLoginRequest{
		Username: "budi",
		Password: "rahasia",
	}

	userResponse, err := AuthService.Login(Ctx, userLoginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, userResponse)
}
