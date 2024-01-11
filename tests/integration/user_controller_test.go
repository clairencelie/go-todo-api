package integration

import (
	"encoding/json"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/model/response"
	"io"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeUserController(t *testing.T) {
	userController := controller.NewUserController(UserService)

	assert.NotNil(t, userController)
}

func TestControllerCreateUser(t *testing.T) {
	ResetDB()

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
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, 201, result.StatusCode)

}

func TestControllerGetUserById(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/user/"+strconv.Itoa(int(userLastInsertId)), nil)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	user := standardResposne.Data.(map[string]any)

	assert.Equal(t, float64(userLastInsertId), user["id"].(float64))
}

func TestControllerGetAllUser(t *testing.T) {
	ResetDB()

	InsertManyUser(TestDb, 5)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/users", nil)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	userList := standardResposne.Data.([]any)

	assert.Len(t, userList, 5)
}

func TestControllerUpdateUser(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	jsonRequest := strings.NewReader(`{
		"username": "budiman",
		"name": "Budiman",
		"email": "budiman@example.xyz",
		"phone_number": "0512345"
	}`)

	request := httptest.NewRequest("PUT", "http://localhost:8080/api/user/"+strconv.Itoa(int(userLastInsertId)), jsonRequest)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestControllerRemoveUser(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/user/"+strconv.Itoa(int(userLastInsertId)), nil)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}
