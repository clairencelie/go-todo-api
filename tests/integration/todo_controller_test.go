package integration

import (
	"encoding/json"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/model/response"
	"io"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeTodoController(t *testing.T) {
	todoController := controller.NewTodoController(TodoService)

	assert.NotNil(t, todoController)
}

func TestControllerCreateTodo(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	todoCreateRequest := request.TodoCreateRequest{
		UserId:      int(userLastInsertId),
		Title:       "Create Todo Test",
		Description: "Create todo test from todo controller",
	}

	jsonTodoCreateRequest, errMarshal := json.Marshal(todoCreateRequest)

	assert.Nil(t, errMarshal)

	requestBody := strings.NewReader(string(jsonTodoCreateRequest))

	request := httptest.NewRequest("POST", "http://localhost:8080/api/todo", requestBody)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, 201, result.StatusCode)
}

func TestControllerGetTodoById(t *testing.T) {
	ResetDB()

	todoLastInsertId := InsertSingleTodo(TestDb)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/todo/"+strconv.Itoa(int(todoLastInsertId)), nil)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	todo := standardResposne.Data.(map[string]any)

	assert.Equal(t, float64(todoLastInsertId), todo["id"].(float64))
}

func TestControllerGetAllTodo(t *testing.T) {
	ResetDB()

	InsertManyTodo(TestDb, 5)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/todos", nil)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	todos := standardResposne.Data.([]any)

	assert.Greater(t, len(todos), 0)
	assert.Len(t, todos, 5)
}

func TestControllerUpdateTodo(t *testing.T) {
	ResetDB()

	todoLastInsertId := InsertSingleTodo(TestDb)

	requestBody := strings.NewReader(`{
		"title": "Update Todo Test",
		"description": "Update todo test from todo controller",
		"is_done": true
	}`)

	request := httptest.NewRequest("PUT", "http://localhost:8080/api/todo/"+strconv.Itoa(int(todoLastInsertId)), requestBody)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestControllerUpdateTodoCompletion(t *testing.T) {
	ResetDB()

	todoLastInsertId := InsertSingleTodo(TestDb)

	request := httptest.NewRequest("PATCH", "http://localhost:8080/api/todo/completion/"+strconv.Itoa(int(todoLastInsertId)), nil)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestControllerRemoveTodo(t *testing.T) {
	ResetDB()

	todoLastInsertId := InsertSingleTodo(TestDb)

	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/todo/"+strconv.Itoa(int(todoLastInsertId)), nil)
	recorder := httptest.NewRecorder()

	Router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}
