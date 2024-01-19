package unit

import (
	"context"
	"database/sql"
	"go_todo_api/internal/model/entity"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (mock *TodoRepositoryMock) Get(ctx context.Context, db *sql.DB, todoId int) (entity.Todo, error) {
	args := mock.Called(ctx, db, todoId)

	if args.Get(1) != nil {
		return args.Get(0).(entity.Todo), args.Get(1).(error)
	}

	return args.Get(0).(entity.Todo), nil
}

func (mock *TodoRepositoryMock) GetUserTodos(ctx context.Context, db *sql.DB, userId int) ([]entity.Todo, error) {
	args := mock.Called(ctx, db, userId)

	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).([]entity.Todo), nil
}

func (mock *TodoRepositoryMock) Insert(ctx context.Context, db *sql.DB, todo request.TodoCreateRequest) error {
	args := mock.Called(ctx, db, todo)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (mock *TodoRepositoryMock) Update(ctx context.Context, db *sql.DB, todo request.TodoUpdateRequest) error {
	args := mock.Called(ctx, db, todo)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (mock *TodoRepositoryMock) UpdateTodoCompletion(ctx context.Context, db *sql.DB, todoId int) error {
	args := mock.Called(ctx, db, todoId)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (mock *TodoRepositoryMock) Delete(ctx context.Context, db *sql.DB, todoId int) error {
	args := mock.Called(ctx, db, todoId)

	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

var todoRepositoryMock = new(TodoRepositoryMock)
var validatorMock = new(ValidatorMock)

func TestTodoServiceFind(t *testing.T) {
	db, _, errDBMock := sqlmock.New()
	assert.NoError(t, errDBMock)

	defer db.Close()

	todoService := service.NewTodoService(db, todoRepositoryMock, validatorMock)

	ctx := context.Background()
	expectedTodo := entity.Todo{
		Id:          1,
		UserId:      1,
		Title:       "Todo Title",
		Description: "Todo description",
		IsDone:      false,
		CreatedAt:   "2023-11-11 11:11:11",
		UpdatedAt:   "2023-11-11 11:11:11",
	}

	todoRepositoryMock.On("Get", ctx, db, 1).Return(expectedTodo, nil)

	todoResponse, errFindTodo := todoService.Find(ctx, 1)
	assert.NoError(t, errFindTodo)

	assert.Equal(t, expectedTodo.Id, todoResponse.Id)
	assert.Equal(t, expectedTodo.UserId, todoResponse.UserId)
	assert.Equal(t, expectedTodo.Title, todoResponse.Title)
	assert.Equal(t, expectedTodo.Description, todoResponse.Description)
	assert.Equal(t, expectedTodo.IsDone, todoResponse.IsDone)
	assert.Equal(t, expectedTodo.CreatedAt, todoResponse.CreatedAt)
	assert.Equal(t, expectedTodo.UpdatedAt, todoResponse.UpdatedAt)
}

func TestTodoServiceFindUserTodos(t *testing.T) {
	db, _, errDBMock := sqlmock.New()
	assert.NoError(t, errDBMock)

	defer db.Close()

	todoService := service.NewTodoService(db, todoRepositoryMock, validatorMock)

	ctx := context.Background()
	expectedTodos := []entity.Todo{}

	for i := 1; i <= 5; i++ {
		todo := entity.Todo{
			Id:          i,
			UserId:      1,
			Title:       "Todo Title " + strconv.Itoa(i),
			Description: "Todo description " + strconv.Itoa(i),
			IsDone:      false,
			CreatedAt:   "2023-11-11 11:11:11",
			UpdatedAt:   "2023-11-11 11:11:11",
		}

		expectedTodos = append(expectedTodos, todo)
	}

	todoRepositoryMock.On("GetUserTodos", ctx, db, 1).Return(expectedTodos, nil)

	todoResponses, errFindUserTodo := todoService.FindUserTodos(ctx, 1)
	assert.NoError(t, errFindUserTodo)

	assert.Equal(t, len(expectedTodos), len(todoResponses))
	assert.Len(t, todoResponses, 5)
}

func TestTodoServiceCreate(t *testing.T) {
	db, _, errDBMock := sqlmock.New()
	assert.NoError(t, errDBMock)

	defer db.Close()

	todoService := service.NewTodoService(db, todoRepositoryMock, validatorMock)

	ctx := context.Background()
	todo := request.TodoCreateRequest{
		UserId:      1,
		Title:       "Todo Title",
		Description: "Todo description",
	}

	validatorMock.On("StructCtx", ctx, todo).Return(nil)
	todoRepositoryMock.On("Insert", ctx, db, todo).Return(nil)

	errCreateTodo := todoService.Create(ctx, todo)
	assert.NoError(t, errCreateTodo)
}

func TestTodoServiceUpdate(t *testing.T) {
	db, _, errDBMock := sqlmock.New()
	assert.NoError(t, errDBMock)

	defer db.Close()

	todoService := service.NewTodoService(db, todoRepositoryMock, validatorMock)

	ctx := context.Background()
	todo := request.TodoUpdateRequest{
		Id:          1,
		Title:       "Todo Title Update",
		Description: "Todo description update",
		IsDone:      true,
	}

	validatorMock.On("StructCtx", ctx, todo).Return(nil)
	todoRepositoryMock.On("Update", ctx, db, todo).Return(nil)

	errUpdateTodo := todoService.Update(ctx, todo)
	assert.NoError(t, errUpdateTodo)
}

func TestTodoServiceUpdateTodoCompletion(t *testing.T) {
	db, _, errDBMock := sqlmock.New()
	assert.NoError(t, errDBMock)

	defer db.Close()

	todoService := service.NewTodoService(db, todoRepositoryMock, validatorMock)

	ctx := context.Background()
	todoRepositoryMock.On("UpdateTodoCompletion", ctx, db, 1).Return(nil)

	errUpdateTodo := todoService.UpdateTodoCompletion(ctx, 1)
	assert.NoError(t, errUpdateTodo)
}

func TestTodoServiceRemove(t *testing.T) {
	db, _, errDBMock := sqlmock.New()
	assert.NoError(t, errDBMock)

	defer db.Close()

	todoService := service.NewTodoService(db, todoRepositoryMock, validatorMock)

	ctx := context.Background()
	todoRepositoryMock.On("Delete", ctx, db, 1).Return(nil)

	errDeleteTodo := todoService.Remove(ctx, 1)
	assert.NoError(t, errDeleteTodo)
}
