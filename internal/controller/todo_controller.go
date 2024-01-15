package controller

import (
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetUserTodos(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateTodoCompletion(w http.ResponseWriter, r *http.Request, params httprouter.Params)
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
	todoCreateRequest := request.TodoCreateRequest{}

	errReadBody := helper.ReadRequestBody(r, &todoCreateRequest)

	if errReadBody != nil {
		helper.WriteErrorResponse(w, errReadBody)
		return
	}

	err := todoController.todoService.Create(r.Context(), todoCreateRequest)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    "new todo created",
	}

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoIdString := params.ByName("todoId")

	todoId, errCastToInt := strconv.Atoi(todoIdString)

	if errCastToInt != nil {
		helper.WriteErrorResponse(w, errCastToInt)
		return
	}

	todoResponse, err := todoController.todoService.Find(r.Context(), todoId)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "todo found",
		Data:       todoResponse,
	}

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) GetUserTodos(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIdString := params.ByName("userId")

	userId, errCastToInt := strconv.Atoi(userIdString)

	if errCastToInt != nil {
		helper.WriteErrorResponse(w, errCastToInt)
		return
	}

	todoResponses, err := todoController.todoService.FindUserTodos(r.Context(), userId)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "todos found",
		Data:       todoResponses,
	}

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoIdString := params.ByName("todoId")

	todoId, errCastToInt := strconv.Atoi(todoIdString)

	if errCastToInt != nil {
		helper.WriteErrorResponse(w, errCastToInt)
		return
	}

	todoUpdateRequest := request.TodoUpdateRequest{
		Id: todoId,
	}

	if errReadBody := helper.ReadRequestBody(r, &todoUpdateRequest); errReadBody != nil {
		helper.WriteErrorResponse(w, errReadBody)
		return
	}

	err := todoController.todoService.Update(r.Context(), todoUpdateRequest)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{StatusCode: http.StatusNoContent}

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) UpdateTodoCompletion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoIdString := params.ByName("todoId")

	todoId, errCastToInt := strconv.Atoi(todoIdString)

	if errCastToInt != nil {
		helper.WriteErrorResponse(w, errCastToInt)
		return
	}

	err := todoController.todoService.UpdateTodoCompletion(r.Context(), todoId)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{StatusCode: http.StatusNoContent}

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoIdString := params.ByName("todoId")

	todoId, errCastToInt := strconv.Atoi(todoIdString)

	if errCastToInt != nil {
		helper.WriteErrorResponse(w, errCastToInt)
		return
	}

	err := todoController.todoService.Remove(r.Context(), todoId)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{StatusCode: http.StatusNoContent}

	helper.WriteResponse(w, responseData)
}
