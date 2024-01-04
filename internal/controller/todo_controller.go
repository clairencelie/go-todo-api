package controller

import (
	"go_todo_api/internal/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type TodoControllerImpl struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		todoService: todoService,
	}
}

func (todoController *TodoControllerImpl) CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (todoController *TodoControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (todoController *TodoControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (todoController *TodoControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (todoController *TodoControllerImpl) Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
