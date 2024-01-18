package integration

import (
	"encoding/json"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/response"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/service"
	testhelper "go_todo_api/tests/test_helper"
	"io"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUserControllerInitialize(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())
	userController := controller.NewUserController(userService)

	assert.NotNil(t, userController)
}

func TestUserControllerCreate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	requestBody := strings.NewReader(`
	{
		"username": "budi",
    	"password": "rahasia",
    	"name": "Budi",
		"email": "budi@example.xyz",
		"phone_number": "0375812938"
	}
	`)

	request := httptest.NewRequest("POST", "http://localhost:8080/api/user", requestBody)
	params := httprouter.Params{}

	recorder := httptest.NewRecorder()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())
	userController := controller.NewUserController(userService)

	userController.CreateUser(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 201, result.StatusCode)
}

func TestUserControllerGetById(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userLastInsertId := testhelper.InsertSingleUser(db)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/user/"+strconv.Itoa(int(userLastInsertId)), nil)
	params := httprouter.Params{
		{
			Key:   "userId",
			Value: strconv.Itoa(int(userLastInsertId)),
		},
	}

	recorder := httptest.NewRecorder()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())
	userController := controller.NewUserController(userService)

	userController.Get(recorder, request, params)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	user := standardResposne.Data.(map[string]any)

	assert.Equal(t, float64(userLastInsertId), user["id"].(float64))
}

func TestUserControllerUpdate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userLastInsertId := testhelper.InsertSingleUser(db)

	jsonRequest := strings.NewReader(`{
		"username": "budiman",
		"name": "Budiman",
		"email": "budiman@example.xyz",
		"phone_number": "0512345"
	}`)

	request := httptest.NewRequest("PUT", "http://localhost:8080/api/user/"+strconv.Itoa(int(userLastInsertId)), jsonRequest)
	params := httprouter.Params{
		{
			Key:   "userId",
			Value: strconv.Itoa(int(userLastInsertId)),
		},
	}

	recorder := httptest.NewRecorder()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())
	userController := controller.NewUserController(userService)

	userController.Update(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestUserControllerRemove(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	userLastInsertId := testhelper.InsertSingleUser(db)

	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/user/"+strconv.Itoa(int(userLastInsertId)), nil)
	params := httprouter.Params{
		{
			Key:   "userId",
			Value: strconv.Itoa(int(userLastInsertId)),
		},
	}

	recorder := httptest.NewRecorder()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validator.New(), helper.HashFunction())
	userController := controller.NewUserController(userService)

	userController.Remove(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}
