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

type TodoServiceMock struct {
	mock.Mock
}

func (mock *TodoServiceMock) Find(ctx context.Context, todoId int) (response.TodoResponse, error) {
	args := mock.Called(ctx, todoId)

	if args.Get(1) != nil {
		return args.Get(0).(response.TodoResponse), args.Get(1).(error)
	}

	return args.Get(0).(response.TodoResponse), nil
}

func (mock *TodoServiceMock) FindUserTodos(ctx context.Context, userId int) ([]response.TodoResponse, error) {
	args := mock.Called(ctx, userId)

	if args.Get(1) != nil {
		return args.Get(0).([]response.TodoResponse), args.Get(1).(error)
	}

	return args.Get(0).([]response.TodoResponse), nil
}

func (mock *TodoServiceMock) Create(ctx context.Context, todo request.TodoCreateRequest) error {
	args := mock.Called(ctx, todo)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (mock *TodoServiceMock) Update(ctx context.Context, todo request.TodoUpdateRequest) error {
	args := mock.Called(ctx, todo)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (mock *TodoServiceMock) UpdateTodoCompletion(ctx context.Context, todoId int) error {
	args := mock.Called(ctx, todoId)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (mock *TodoServiceMock) Remove(ctx context.Context, todoId int) error {
	args := mock.Called(ctx, todoId)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

var todoServiceMock = new(TodoServiceMock)

func TestTodoControllerCreateTodo(t *testing.T) {
	todoCreateRequest := request.TodoCreateRequest{
		UserId:      1,
		Title:       "Create Todo Test",
		Description: "Create todo test from todo controller",
	}

	jsonTodoCreateRequest, errMarshal := json.Marshal(todoCreateRequest)

	assert.Nil(t, errMarshal)

	requestBody := strings.NewReader(string(jsonTodoCreateRequest))

	request := httptest.NewRequest("POST", "http://localhost:8080/api/todo", requestBody)
	params := httprouter.Params{}

	recorder := httptest.NewRecorder()

	todoController := controller.NewTodoController(todoServiceMock)

	todoServiceMock.On("Create", request.Context(), mock.AnythingOfType("request.TodoCreateRequest")).Return(nil)

	todoController.CreateTodo(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 201, result.StatusCode)
}

func TestTodoControllerGetById(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/api/todo/1", nil)
	params := httprouter.Params{
		{
			Key:   "todoId",
			Value: "1",
		},
	}

	recorder := httptest.NewRecorder()

	todoController := controller.NewTodoController(todoServiceMock)
	todoResponse := response.TodoResponse{
		Id:          1,
		UserId:      1,
		Title:       "Todo Title",
		Description: "Todo description",
		IsDone:      false,
		CreatedAt:   "2024-01-01 11:11:11",
		UpdatedAt:   "2024-01-01 11:11:11",
	}

	todoServiceMock.On("Find", request.Context(), 1).Return(todoResponse, nil)

	todoController.Get(recorder, request, params)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	todo := standardResposne.Data.(map[string]any)

	assert.Equal(t, float64(todoResponse.Id), todo["id"])
	assert.Equal(t, float64(todoResponse.UserId), todo["user_id"])
	assert.Equal(t, todoResponse.Title, todo["title"])
	assert.Equal(t, todoResponse.Description, todo["description"])
	assert.Equal(t, todoResponse.IsDone, todo["is_done"])
	assert.Equal(t, todoResponse.CreatedAt, todo["created_at"])
	assert.Equal(t, todoResponse.UpdatedAt, todo["updated_at"])
}

func TestTodoControllerGetUserTodos(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/api/user/1/todo", nil)
	params := httprouter.Params{
		{
			Key:   "userId",
			Value: "1",
		},
	}

	recorder := httptest.NewRecorder()

	todoController := controller.NewTodoController(todoServiceMock)

	todoResponses := []response.TodoResponse{}

	for i := 1; i <= 5; i++ {
		todoResponse := response.TodoResponse{
			Id:          i,
			UserId:      1,
			Title:       "Todo Title",
			Description: "Todo description",
			IsDone:      false,
			CreatedAt:   "2024-01-01 11:11:11",
			UpdatedAt:   "2024-01-01 11:11:11",
		}

		todoResponses = append(todoResponses, todoResponse)
	}

	todoServiceMock.On("FindUserTodos", request.Context(), 1).Return(todoResponses, nil)

	todoController.GetUserTodos(recorder, request, params)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	todos := standardResposne.Data.([]any)

	assert.Len(t, todos, 5)
}

func TestTodoControllerUpdate(t *testing.T) {
	requestBody := strings.NewReader(`{
		"title": "Update Todo Test",
		"description": "Update todo test from todo controller",
		"is_done": true
	}`)

	request := httptest.NewRequest("PUT", "http://localhost:8080/api/todo/1", requestBody)
	params := httprouter.Params{
		{
			Key:   "todoId",
			Value: "1",
		},
	}

	recorder := httptest.NewRecorder()

	todoController := controller.NewTodoController(todoServiceMock)

	todoServiceMock.On("Update", request.Context(), mock.AnythingOfType("request.TodoUpdateRequest")).Return(nil)

	todoController.Update(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestTodoControllerUpdateTodoCompletion(t *testing.T) {
	request := httptest.NewRequest("PATCH", "http://localhost:8080/api/todo/completion/1", nil)
	params := httprouter.Params{
		{
			Key:   "todoId",
			Value: "1",
		},
	}

	recorder := httptest.NewRecorder()

	todoController := controller.NewTodoController(todoServiceMock)

	todoServiceMock.On("UpdateTodoCompletion", request.Context(), 1).Return(nil)

	todoController.UpdateTodoCompletion(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestTodoControllerRemove(t *testing.T) {
	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/todo/1", nil)
	params := httprouter.Params{
		{
			Key:   "todoId",
			Value: "1",
		},
	}

	recorder := httptest.NewRecorder()

	todoController := controller.NewTodoController(todoServiceMock)

	todoServiceMock.On("Remove", request.Context(), 1).Return(nil)

	todoController.Remove(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}
