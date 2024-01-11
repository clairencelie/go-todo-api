package integration

import (
	"encoding/json"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/model/response"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeAuthController(t *testing.T) {
	authController := controller.NewAuthController(AuthService)

	assert.NotNil(t, authController)
}

func TestControllerLogin(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	jsonRequest := strings.NewReader(`{
		"username": "budi",
		"password": "rahasia"
	}`)

	request := httptest.NewRequest("POST", "http://localhost:8080/api/login", jsonRequest)
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
