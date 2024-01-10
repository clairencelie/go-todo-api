package integration

import (
	"context"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodoRepository(t *testing.T) {
	todoRepository := repository.NewTodoRepository()

	assert.NotNil(t, todoRepository)
}

func TestGetTodoById(t *testing.T) {
	ResetDB()

	todoLastInsertId := InsertSingleTodo(TestDb)

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	todo, err := todoRepository.Get(ctx, TestDb, int(todoLastInsertId))

	assert.Nil(t, err)
	assert.NotNil(t, todo)

	assert.Equal(t, todoLastInsertId, int64(todo.Id))
	assert.Equal(t, "todo 1", todo.Title)
	assert.Equal(t, "deskripsi todo 1", todo.Description)
	assert.False(t, todo.IsDone)
}

func TestGetAllTodo(t *testing.T) {
	ResetDB()

	InsertManyTodo(TestDb, 5)

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	todos, err := todoRepository.GetAll(ctx, TestDb)

	assert.Nil(t, err)
	assert.NotNil(t, todos)

	assert.Greater(t, len(todos), 0)
	assert.Len(t, todos, 5)
}

func TestInsertTodo(t *testing.T) {
	ResetDB()

	userLastInsertId := InsertSingleUser(TestDb)

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	tx, errTx := TestDb.Begin()

	assert.Nil(t, errTx)

	todoCreateRequest := request.TodoCreateRequest{
		UserId:      int(userLastInsertId),
		Title:       "todo test",
		Description: "todo single insertion test",
	}

	err := todoRepository.Insert(ctx, tx, todoCreateRequest)

	assert.Nil(t, err)
	if err == nil {
		tx.Commit()
	}
}

func TestUpdateTodo(t *testing.T) {
	ResetDB()

	todoLastInsertId := InsertSingleTodo(TestDb)

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	tx, errTx := TestDb.Begin()

	assert.Nil(t, errTx)

	todoUpdateRequest := request.TodoUpdateRequest{
		Id:          int(todoLastInsertId),
		Title:       "test update todo",
		Description: "test update todo description",
		IsDone:      true,
	}

	err := todoRepository.Update(ctx, tx, todoUpdateRequest)

	assert.Nil(t, err)
	if err == nil {
		tx.Commit()
	}
}

func TestDeleteTodo(t *testing.T) {
	ResetDB()

	todoLastInsertId := InsertSingleTodo(TestDb)

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	tx, errTx := TestDb.Begin()

	assert.Nil(t, errTx)

	err := todoRepository.Delete(ctx, tx, int(todoLastInsertId))

	assert.Nil(t, err)
	if err == nil {
		tx.Commit()
	}
}
