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

type UserServiceMock struct {
	mock.Mock
}

func (mock *UserServiceMock) Find(ctx context.Context, userId int) (response.UserResponse, error) {
	args := mock.Called(ctx, userId)

	if args.Get(1) != nil {
		return args.Get(0).(response.UserResponse), args.Get(1).(error)
	}

	return args.Get(0).(response.UserResponse), nil
}

func (mock *UserServiceMock) Create(ctx context.Context, user request.UserCreateRequest) error {
	args := mock.Called(ctx, user)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (mock *UserServiceMock) Update(ctx context.Context, user request.UserUpdateRequest) error {
	args := mock.Called(ctx, user)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (mock *UserServiceMock) Remove(ctx context.Context, userId int) error {
	args := mock.Called(ctx, userId)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func TestUserControllerCreate(t *testing.T) {
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

	userServiceMock := new(UserServiceMock)
	userController := controller.NewUserController(userServiceMock)

	userServiceMock.On("Create", request.Context(), mock.AnythingOfType("request.UserCreateRequest")).Return(nil)

	userController.CreateUser(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 201, result.StatusCode)
}

func TestUserControllerGetById(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/api/user/1", nil)
	params := httprouter.Params{
		{
			Key:   "userId",
			Value: "1",
		},
	}

	recorder := httptest.NewRecorder()

	userServiceMock := new(UserServiceMock)
	userController := controller.NewUserController(userServiceMock)

	userResponse := response.UserResponse{
		Id:          1,
		Username:    "apollo",
		Name:        "Apollo",
		Email:       "apollo@example.xyz",
		PhoneNumber: "08275819224",
		CreatedAt:   "2023-11-11 11:11:11",
	}

	userServiceMock.On("Find", request.Context(), 1).Return(userResponse, nil)

	userController.Get(recorder, request, params)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	user := standardResposne.Data.(map[string]any)

	assert.Equal(t, float64(1), user["id"])
	assert.Equal(t, userResponse.Username, user["username"])
	assert.Equal(t, userResponse.Name, user["name"])
	assert.Equal(t, userResponse.Email, user["email"])
	assert.Equal(t, userResponse.PhoneNumber, user["phone_number"])
	assert.Equal(t, userResponse.CreatedAt, user["created_at"])
}

func TestUserControllerUpdate(t *testing.T) {
	jsonRequest := strings.NewReader(`{
		"username": "budiman",
		"name": "Budiman",
		"email": "budiman@example.xyz",
		"phone_number": "0512345"
	}`)

	request := httptest.NewRequest("PUT", "http://localhost:8080/api/user/1", jsonRequest)
	params := httprouter.Params{
		{
			Key:   "userId",
			Value: "1",
		},
	}

	recorder := httptest.NewRecorder()

	userServiceMock := new(UserServiceMock)
	userController := controller.NewUserController(userServiceMock)

	userServiceMock.On("Update", request.Context(), mock.AnythingOfType("request.UserUpdateRequest")).Return(nil)

	userController.Update(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestUserControllerDelete(t *testing.T) {
	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/user/1", nil)
	params := httprouter.Params{
		{
			Key:   "userId",
			Value: "1",
		},
	}

	recorder := httptest.NewRecorder()

	userServiceMock := new(UserServiceMock)
	userController := controller.NewUserController(userServiceMock)

	userServiceMock.On("Remove", request.Context(), 1).Return(nil)

	userController.Remove(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}
