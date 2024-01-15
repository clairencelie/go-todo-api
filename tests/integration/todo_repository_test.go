package integration

import (
	"context"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	testhelper "go_todo_api/tests/test_helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoRepositoryInitialize(t *testing.T) {
	todoRepository := repository.NewTodoRepository()

	assert.NotNil(t, todoRepository)
}

func TestTodoRepositoryGetById(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInsertId := testhelper.InsertSingleTodo(db)

	todoRepository := repository.NewTodoRepository()

	todo, err := todoRepository.Get(context.Background(), db, int(todoLastInsertId))

	assert.Nil(t, err)
	assert.NotNil(t, todo)

	assert.Equal(t, todoLastInsertId, int64(todo.Id))
	assert.Equal(t, "todo 1", todo.Title)
	assert.Equal(t, "deskripsi todo 1", todo.Description)
	assert.False(t, todo.IsDone)
}

func TestTodoRepositoryInsert(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	userLastInsertId := testhelper.InsertSingleUser(db)

	todoRepository := repository.NewTodoRepository()

	todoCreateRequest := request.TodoCreateRequest{
		UserId:      int(userLastInsertId),
		Title:       "todo test",
		Description: "todo single insertion test",
	}

	err := todoRepository.Insert(context.Background(), db, todoCreateRequest)

	assert.Nil(t, err)
}

func TestTodoRepositoryUpdate(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInsertId := testhelper.InsertSingleTodo(db)

	todoRepository := repository.NewTodoRepository()

	todoUpdateRequest := request.TodoUpdateRequest{
		Id:          int(todoLastInsertId),
		Title:       "test update todo",
		Description: "test update todo description",
		IsDone:      true,
	}

	err := todoRepository.Update(context.Background(), db, todoUpdateRequest)

	assert.Nil(t, err)
}

func TestTodoRepositoryDelete(t *testing.T) {
	db, errDbConn := setupDb()

	assert.Nil(t, errDbConn)

	defer db.Close()

	todoLastInsertId := testhelper.InsertSingleTodo(db)

	todoRepository := repository.NewTodoRepository()

	err := todoRepository.Delete(context.Background(), db, int(todoLastInsertId))

	assert.Nil(t, err)
}
