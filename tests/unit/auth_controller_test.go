package unit

import (
	"context"
	"encoding/json"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/model/response"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AuthServiceMock struct {
	mock.Mock
}

func (mock *AuthServiceMock) Login(ctx context.Context, loginRequest request.UserLoginRequest) (response.UserResponse, error) {
	args := mock.Called(ctx, loginRequest)

	if args.Get(1) != nil {
		return args.Get(0).(response.UserResponse), args.Get(1).(error)
	}

	return args.Get(0).(response.UserResponse), nil
}

func TestAuthControllerLogin(t *testing.T) {
	jsonRequest := strings.NewReader(`{
		"username": "apollo",
		"password": "secret"
	}`)

	request := httptest.NewRequest("POST", "http://localhost:8080/api/login", jsonRequest)
	params := httprouter.Params{}

	recorder := httptest.NewRecorder()

	authServiceMock := new(AuthServiceMock)
	authController := controller.NewAuthController(authServiceMock)

	userResponse := response.UserResponse{
		Id:          1,
		Username:    "apollo",
		Name:        "Apollo",
		Email:       "apollo@example.xyz",
		PhoneNumber: "08275819224",
		CreatedAt:   "2023-11-11 11:11:11",
	}

	authServiceMock.On("Login", request.Context(), mock.AnythingOfType("request.UserLoginRequest")).Return(userResponse, nil)

	authController.Login(recorder, request, params)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	user := standardResposne.Data.(map[string]any)

	assert.Equal(t, float64(userResponse.Id), user["id"].(float64))
	assert.Equal(t, userResponse.Username, user["username"])
	assert.Equal(t, userResponse.Name, user["name"])
	assert.Equal(t, userResponse.Email, user["email"])
	assert.Equal(t, userResponse.PhoneNumber, user["phone_number"])
	assert.Equal(t, userResponse.CreatedAt, user["created_at"])
}
