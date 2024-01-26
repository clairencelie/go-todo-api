package integration

import (
	"encoding/json"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/model/response"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/service"
	testhelper "go_todo_api/tests/test_helper"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestAuthControllerInitialize(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(db, userRepository, validator.New())
	authController := controller.NewAuthController(authService)

	assert.NotNil(t, authController)
}

func TestAuthControllerLogin(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userLastInsertId := testhelper.InsertSingleUser(db)

	jsonRequest := strings.NewReader(`{
		"username": "budi",
		"password": "rahasia"
	}`)

	request := httptest.NewRequest("POST", "http://localhost:8080/api/login", jsonRequest)
	recorder := httptest.NewRecorder()

	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(db, userRepository, validator.New())
	authController := controller.NewAuthController(authService)

	params := httprouter.Params{}

	authController.Login(recorder, request, params)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.NoError(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	user := standardResposne.Data.(map[string]any)

	assert.Equal(t, float64(userLastInsertId), user["id"].(float64))
}
