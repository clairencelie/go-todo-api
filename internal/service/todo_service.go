package service

import (
	"context"
	"database/sql"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/model/response"
	"go_todo_api/internal/repository"
	customvalidator "go_todo_api/internal/validator"
)

type TodoService interface {
	Find(ctx context.Context, todoId int) (response.TodoResponse, error)
	FindUserTodos(ctx context.Context, userId int) ([]response.TodoResponse, error)
	Create(ctx context.Context, todo request.TodoCreateRequest) error
	Update(ctx context.Context, todo request.TodoUpdateRequest) error
	UpdateTodoCompletion(ctx context.Context, todoId int) error
	Remove(ctx context.Context, todoId int) error
}

type TodoServiceImpl struct {
	db             *sql.DB
	todoRepository repository.TodoRepository
	validate       customvalidator.CustomValidator
}

func NewTodoService(db *sql.DB, todoRepository repository.TodoRepository, validate customvalidator.CustomValidator) TodoService {
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

func (todoService *TodoServiceImpl) FindUserTodos(ctx context.Context, userId int) ([]response.TodoResponse, error) {
	todos, err := todoService.todoRepository.GetUserTodos(ctx, todoService.db, userId)

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

	err := todoService.todoRepository.Insert(ctx, todoService.db, todo)

	if err != nil {
		return err
	}

	return nil
}

func (todoService *TodoServiceImpl) Update(ctx context.Context, todo request.TodoUpdateRequest) error {
	errValidation := todoService.validate.StructCtx(ctx, todo)

	if errValidation != nil {
		return errValidation
	}

	err := todoService.todoRepository.Update(ctx, todoService.db, todo)

	if err != nil {
		return err
	}

	return nil
}

func (todoService *TodoServiceImpl) UpdateTodoCompletion(ctx context.Context, todoId int) error {
	err := todoService.todoRepository.UpdateTodoCompletion(ctx, todoService.db, todoId)

	if err != nil {
		return err
	}

	return nil
}

func (todoService *TodoServiceImpl) Remove(ctx context.Context, todoId int) error {
	err := todoService.todoRepository.Delete(ctx, todoService.db, todoId)

	if err != nil {
		return err
	}

	return nil
}
