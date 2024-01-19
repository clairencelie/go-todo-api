package unit

import (
	"context"
	"database/sql/driver"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var todoRepository = repository.NewTodoRepository()

func TestTodoRepositoryGet(t *testing.T) {
	db, mock, errDBMock := sqlmock.New()

	assert.NoError(t, errDBMock)

	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "user_id", "title", "description", "is_done", "created_at", "updated_at"}).AddRow(1, 1, "Todo Title", "Todo description", false, "2024-01-01", "2024-01-01")

	mock.ExpectPrepare("SELECT id, user_id, title, description, is_done, created_at, updated_at FROM todos").ExpectQuery().WithArgs(1).WillReturnRows(row)

	todo, errGetTodo := todoRepository.Get(context.Background(), db, 1)

	assert.NoError(t, errGetTodo)
	assert.Equal(t, 1, todo.Id)
	assert.Equal(t, 1, todo.UserId)
	assert.Equal(t, "Todo Title", todo.Title)
	assert.Equal(t, "Todo description", todo.Description)
	assert.False(t, todo.IsDone)
	assert.Equal(t, "2024-01-01", todo.CreatedAt)
	assert.Equal(t, "2024-01-01", todo.UpdatedAt)

	errMockExpectations := mock.ExpectationsWereMet()
	assert.NoError(t, errMockExpectations)
}

func TestTodoRepositoryGetUserTodos(t *testing.T) {
	db, mock, errDBMock := sqlmock.New()

	assert.NoError(t, errDBMock)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "user_id", "title", "description", "is_done", "created_at", "updated_at"})

	for i := 1; i <= 3; i++ {
		value := []driver.Value{i, 1, "Todo Title " + strconv.Itoa(i), "Todo description" + strconv.Itoa(i), false, "2024-01-01", "2024-01-01"}
		rows.AddRows(value)
	}

	mock.ExpectPrepare("SELECT id, user_id, title, description, is_done, created_at, updated_at FROM todos").ExpectQuery().WithArgs(1).WillReturnRows(rows)

	todos, errGetTodo := todoRepository.GetUserTodos(context.Background(), db, 1)

	assert.NoError(t, errGetTodo)
	assert.Len(t, todos, 3)

	errMockExpectations := mock.ExpectationsWereMet()
	assert.NoError(t, errMockExpectations)
}

func TestTodoRepositoryInsert(t *testing.T) {
	db, mock, errDBMock := sqlmock.New()

	assert.NoError(t, errDBMock)

	defer db.Close()

	todo := request.TodoCreateRequest{
		UserId:      1,
		Title:       "Todo Title",
		Description: "Todo description",
	}

	mock.ExpectPrepare("INSERT INTO todos").ExpectExec().WithArgs(todo.UserId, todo.Title, todo.Description).WillReturnResult(sqlmock.NewResult(1, 1))

	errInsertTodo := todoRepository.Insert(context.Background(), db, todo)
	assert.NoError(t, errInsertTodo)

	errMockExpectations := mock.ExpectationsWereMet()
	assert.NoError(t, errMockExpectations)
}

func TestTodoRepositoryUpdate(t *testing.T) {
	db, mock, errDBMock := sqlmock.New()

	assert.NoError(t, errDBMock)

	defer db.Close()

	todoUpdate := request.TodoUpdateRequest{
		Id:          1,
		Title:       "Update Todo Title",
		Description: "Update todo description",
		IsDone:      true,
	}

	mock.ExpectPrepare("UPDATE todos SET").ExpectExec().WithArgs(todoUpdate.Title, todoUpdate.Description, todoUpdate.IsDone, todoUpdate.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	errUpdateTodo := todoRepository.Update(context.Background(), db, todoUpdate)
	assert.NoError(t, errUpdateTodo)

	errMockExpectations := mock.ExpectationsWereMet()
	assert.NoError(t, errMockExpectations)
}

func TestTodoRepositoryUpdateCompletion(t *testing.T) {
	db, mock, errDBMock := sqlmock.New()

	assert.NoError(t, errDBMock)

	defer db.Close()

	mock.ExpectPrepare("UPDATE todos SET").ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	errUpdateTodoCompletion := todoRepository.UpdateTodoCompletion(context.Background(), db, 1)
	assert.NoError(t, errUpdateTodoCompletion)

	errMockExpectations := mock.ExpectationsWereMet()
	assert.NoError(t, errMockExpectations)
}

func TestTodoRepositoryDelete(t *testing.T) {
	db, mock, errDBMock := sqlmock.New()

	assert.NoError(t, errDBMock)

	defer db.Close()

	mock.ExpectPrepare("DELETE FROM todos").ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	errUpdateTodoCompletion := todoRepository.Delete(context.Background(), db, 1)
	assert.NoError(t, errUpdateTodoCompletion)

	errMockExpectations := mock.ExpectationsWereMet()
	assert.NoError(t, errMockExpectations)
}
