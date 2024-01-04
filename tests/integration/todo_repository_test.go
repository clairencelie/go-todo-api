package integration

import (
	"context"
	"go_todo_api/database"
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
	db, _ := database.NewDB("./../../")

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	todo, err := todoRepository.Get(ctx, db, 1)

	assert.Nil(t, err)
	assert.NotNil(t, todo)

	assert.Equal(t, 1, todo.Id)
	assert.Equal(t, "Finish Repository Module", todo.Title)
	assert.Equal(t, "Finish repository module in go-todo project today", todo.Description)
	assert.False(t, todo.IsDone)
}

func TestGetAllTodo(t *testing.T) {
	db, _ := database.NewDB("./../../")

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	todos, err := todoRepository.GetAll(ctx, db)

	assert.Nil(t, err)
	assert.NotNil(t, todos)

	assert.Greater(t, len(todos), 0)
}

func TestInsertTodo(t *testing.T) {
	db, _ := database.NewDB("./../../")

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	tx, errTx := db.Begin()

	assert.Nil(t, errTx)

	todoCreateRequest := request.TodoCreateRequest{
		UserId:      2,
		Title:       "Push Leng Mobel Lejen",
		Description: "Push leng sampai mitik glory kacks",
	}

	err := todoRepository.Insert(ctx, tx, todoCreateRequest)

	assert.Nil(t, err)
	if err == nil {
		tx.Commit()
	}
}

func TestUpdateTodo(t *testing.T) {
	db, _ := database.NewDB("./../../")

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	tx, errTx := db.Begin()

	assert.Nil(t, errTx)

	todoUpdateRequest := request.TodoUpdateRequest{
		Id:          2,
		Title:       "Push Leng Mobel Lejen.",
		Description: "Push leng sampai mitik aja kacks",
		IsDone:      true,
	}

	err := todoRepository.Update(ctx, tx, todoUpdateRequest)

	assert.Nil(t, err)
	if err == nil {
		tx.Commit()
	}
}

func TestDeleteTodo(t *testing.T) {
	db, _ := database.NewDB("./../../")

	todoRepository := repository.NewTodoRepository()

	ctx := context.Background()

	tx, errTx := db.Begin()

	assert.Nil(t, errTx)

	err := todoRepository.Delete(ctx, tx, 2)

	assert.Nil(t, err)
	if err == nil {
		tx.Commit()
	}
}
