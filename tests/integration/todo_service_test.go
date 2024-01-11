package integration

import (
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeTodoService(t *testing.T) {
	todoService := service.NewTodoService(TestDb, TodoRepository, Validator)

	assert.NotNil(t, todoService)
}

func TestServiceCreateTodo(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	todoCreateRequest := request.TodoCreateRequest{
		UserId:      int(userLastInsertId),
		Title:       "todo test",
		Description: "create todo from todo service",
	}

	err := TodoService.Create(Ctx, todoCreateRequest)

	assert.Nil(t, err)
}

func TestServiceFindTodoById(t *testing.T) {
	ResetDB()

	todoLastInsertid := InsertSingleTodo(TestDb)

	todoResponse, err := TodoService.Find(Ctx, int(todoLastInsertid))

	assert.Nil(t, err)
	assert.NotNil(t, todoResponse)
}

func TestServiceFindAllTodo(t *testing.T) {
	ResetDB()

	InsertManyTodo(TestDb, 5)

	todoResponses, err := TodoService.FindAll(Ctx)

	assert.Nil(t, err)
	assert.Greater(t, len(todoResponses), 0)
	assert.Len(t, todoResponses, 5)
}

func TestServiceUpdateTodo(t *testing.T) {
	ResetDB()

	todoLastInserId := InsertSingleTodo(TestDb)

	todoUpdateRequest := request.TodoUpdateRequest{
		Id:          int(todoLastInserId),
		Title:       "Update Todo",
		Description: "Update todo from todo service",
		IsDone:      true,
	}

	err := TodoService.Update(Ctx, todoUpdateRequest)

	assert.Nil(t, err)
}

func TestServiceUpdateTodoCompletion(t *testing.T) {
	ResetDB()

	todoLastInserId := InsertSingleTodo(TestDb)

	err := TodoService.UpdateTodoCompletion(Ctx, int(todoLastInserId))

	assert.Nil(t, err)
}

func TestServiceRemoveTodo(t *testing.T) {
	ResetDB()

	todoLastInserId := InsertSingleTodo(TestDb)

	err := TodoService.Remove(Ctx, int(todoLastInserId))

	assert.Nil(t, err)
}
