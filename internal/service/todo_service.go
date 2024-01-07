package service

import (
	"context"
	"database/sql"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/model/response"
	"go_todo_api/internal/repository"

	"github.com/go-playground/validator/v10"
)

type TodoService interface {
	Find(ctx context.Context, todoId int) (response.TodoResponse, error)
	FindAll(ctx context.Context) ([]response.TodoResponse, error)
	Create(ctx context.Context, todo request.TodoCreateRequest) error
	Update(ctx context.Context, todo request.TodoUpdateRequest) error
	Remove(ctx context.Context, todoId int) error
}

type TodoServiceImpl struct {
	db             *sql.DB
	todoRepository repository.TodoRepository
	validate       *validator.Validate
}

func NewTodoService(db *sql.DB, todoRepository repository.TodoRepository, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		db:             db,
		todoRepository: todoRepository,
		validate:       validate,
	}
}

func (todoService *TodoServiceImpl) Find(ctx context.Context, todoId int) (response.TodoResponse, error) {
	todo, err := todoService.todoRepository.Get(ctx, todoService.db, todoId)

	if err != nil {
		return response.TodoResponse{}, err
	}

	todoResponse := response.TodoResponse{
		Id:          todo.Id,
		UserId:      todo.UserId,
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}

	return todoResponse, nil

}

func (todoService *TodoServiceImpl) FindAll(ctx context.Context) ([]response.TodoResponse, error) {
	todos, err := todoService.todoRepository.GetAll(ctx, todoService.db)

	if err != nil {
		return nil, err
	}

	todoResponses := []response.TodoResponse{}

	for _, todo := range todos {
		todoResponse := response.TodoResponse{
			Id:          todo.Id,
			UserId:      todo.UserId,
			Title:       todo.Title,
			Description: todo.Description,
			IsDone:      todo.IsDone,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		}

		todoResponses = append(todoResponses, todoResponse)
	}

	return todoResponses, nil
}

func (todoService *TodoServiceImpl) Create(ctx context.Context, todo request.TodoCreateRequest) error {
	errValidation := todoService.validate.StructCtx(ctx, todo)
	if errValidation != nil {
		return errValidation
	}

	tx, errTxBegin := todoService.db.Begin()

	if errTxBegin != nil {
		return errTxBegin
	}

	err := todoService.todoRepository.Insert(ctx, tx, todo)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (todoService *TodoServiceImpl) Update(ctx context.Context, todo request.TodoUpdateRequest) error {
	errValidation := todoService.validate.StructCtx(ctx, todo)
	if errValidation != nil {
		return errValidation
	}

	tx, errTxBegin := todoService.db.Begin()

	if errTxBegin != nil {
		return errTxBegin
	}

	err := todoService.todoRepository.Update(ctx, tx, todo)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (todoService *TodoServiceImpl) Remove(ctx context.Context, todoId int) error {
	tx, errTxBegin := todoService.db.Begin()

	if errTxBegin != nil {
		return errTxBegin
	}

	err := todoService.todoRepository.Delete(ctx, tx, todoId)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
