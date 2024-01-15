package integration

import (
	"encoding/json"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/model/request"
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

func TestTodoControllerInitialize(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())
	todoController := controller.NewTodoController(todoService)

	assert.NotNil(t, todoController)
}

func TestTodoControllerCreate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userLastInsertId := testhelper.InsertSingleUser(db)

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

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())
	todoController := controller.NewTodoController(todoService)

	params := httprouter.Params{}

	todoController.CreateTodo(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 201, result.StatusCode)
}

func TestTodoControllerGetById(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInsertId := testhelper.InsertSingleTodo(db)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/todo/"+strconv.Itoa(int(todoLastInsertId)), nil)
	recorder := httptest.NewRecorder()

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())
	todoController := controller.NewTodoController(todoService)

	params := httprouter.Params{
		{
			Key:   "todoId",
			Value: strconv.Itoa(int(todoLastInsertId)),
		},
	}

	todoController.Get(recorder, request, params)

	result := recorder.Result()
	bytes, err := io.ReadAll(result.Body)

	assert.Equal(t, 200, result.StatusCode)
	assert.Nil(t, err)

	standardResposne := response.StandardResponse{}

	json.Unmarshal(bytes, &standardResposne)

	todo := standardResposne.Data.(map[string]any)

	assert.Equal(t, float64(todoLastInsertId), todo["id"].(float64))
}

func TestTodoControllerUpdate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInsertId := testhelper.InsertSingleTodo(db)

	requestBody := strings.NewReader(`{
		"title": "Update Todo Test",
		"description": "Update todo test from todo controller",
		"is_done": true
	}`)

	request := httptest.NewRequest("PUT", "http://localhost:8080/api/todo/"+strconv.Itoa(int(todoLastInsertId)), requestBody)
	recorder := httptest.NewRecorder()

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())
	todoController := controller.NewTodoController(todoService)

	params := httprouter.Params{
		{
			Key:   "todoId",
			Value: strconv.Itoa(int(todoLastInsertId)),
		},
	}

	todoController.Update(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestTodoControllerUpdateTodoCompletion(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInsertId := testhelper.InsertSingleTodo(db)

	request := httptest.NewRequest("PATCH", "http://localhost:8080/api/todo/completion/"+strconv.Itoa(int(todoLastInsertId)), nil)
	recorder := httptest.NewRecorder()

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())
	todoController := controller.NewTodoController(todoService)

	params := httprouter.Params{
		{
			Key:   "todoId",
			Value: strconv.Itoa(int(todoLastInsertId)),
		},
	}

	todoController.UpdateTodoCompletion(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}

func TestTodoControllerRemove(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInsertId := testhelper.InsertSingleTodo(db)

	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/todo/"+strconv.Itoa(int(todoLastInsertId)), nil)
	recorder := httptest.NewRecorder()

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())
	todoController := controller.NewTodoController(todoService)

	params := httprouter.Params{
		{
			Key:   "todoId",
			Value: strconv.Itoa(int(todoLastInsertId)),
		},
	}

	todoController.Remove(recorder, request, params)

	result := recorder.Result()

	assert.Equal(t, 204, result.StatusCode)
}
