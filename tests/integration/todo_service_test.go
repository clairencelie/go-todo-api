package integration

import (
	"context"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/service"
	testhelper "go_todo_api/tests/test_helper"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestTodoServiceInitialize(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())

	assert.NotNil(t, todoService)
}

func TestTodoServiceCreate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userLastInsertId := testhelper.InsertSingleUser(db)

	todoCreateRequest := request.TodoCreateRequest{
		UserId:      int(userLastInsertId),
		Title:       "todo test",
		Description: "create todo from todo service",
	}

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())

	err := todoService.Create(context.Background(), todoCreateRequest)

	assert.Nil(t, err)
}

func TestTodoServiceFindById(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInsertid := testhelper.InsertSingleTodo(db)

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())

	todoResponse, err := todoService.Find(context.Background(), int(todoLastInsertid))

	assert.Nil(t, err)
	assert.NotNil(t, todoResponse)
}

func TestTodoServiceUpdate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInserId := testhelper.InsertSingleTodo(db)

	todoUpdateRequest := request.TodoUpdateRequest{
		Id:          int(todoLastInserId),
		Title:       "Update Todo",
		Description: "Update todo from todo service",
		IsDone:      true,
	}

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())

	err := todoService.Update(context.Background(), todoUpdateRequest)

	assert.Nil(t, err)
}

func TestTodoServiceUpdateTodoCompletion(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInserId := testhelper.InsertSingleTodo(db)

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())

	err := todoService.UpdateTodoCompletion(context.Background(), int(todoLastInserId))

	assert.Nil(t, err)
}

func TestTodoServiceRemove(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInserId := testhelper.InsertSingleTodo(db)

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, validator.New())

	err := todoService.Remove(context.Background(), int(todoLastInserId))

	assert.Nil(t, err)
}
