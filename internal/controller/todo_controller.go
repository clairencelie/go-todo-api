package controller

import (
	"errors"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
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
	todoCreateRequest := request.TodoCreateRequest{}

	errReadBody := helper.ReadRequestBody(r, &todoCreateRequest)

	responseData := helper.ResponseData{}

	if errReadBody != nil {
		responseData.StatusCode = 500
		responseData.Message = "internal server error"
		responseData.Err = errReadBody

		helper.WriteResponse(w, responseData)
		return
	}

	err := todoController.todoService.Create(r.Context(), todoCreateRequest)

	if err != nil {
		if errValidation, ok := err.(validator.ValidationErrors); ok {
			responseData.StatusCode = 400
			responseData.Message = "validation error"
			responseData.Data = errValidation.Error()
		} else {
			responseData.StatusCode = 500
			responseData.Message = "internal server error"
			responseData.Err = err
		}

		helper.WriteResponse(w, responseData)
		return
	}

	responseData.StatusCode = 201
	responseData.Message = "new todo created"

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoIdString := params.ByName("todoId")

	todoId, errCastToInt := strconv.Atoi(todoIdString)

	responseData := helper.ResponseData{}

	if errCastToInt != nil {
		responseData.StatusCode = 500
		responseData.Message = "failed to cast todo id to int"
		responseData.Err = errCastToInt

		helper.WriteResponse(w, responseData)
		return
	}

	todoResponse, err := todoController.todoService.Find(r.Context(), todoId)

	if err != nil {
		if errors.Is(repository.ErrNotFound, err) {
			responseData.StatusCode = 404
			responseData.Message = "todo not found"
		} else {
			responseData.StatusCode = 500
			responseData.Message = "internal server error"
			responseData.Err = err
		}

		helper.WriteResponse(w, responseData)
		return
	}

	responseData.StatusCode = 200
	responseData.Message = "todo found"
	responseData.Data = todoResponse

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseData := helper.ResponseData{}

	todoResponses, err := todoController.todoService.FindAll(r.Context())

	if err != nil {
		responseData.StatusCode = 500
		responseData.Message = "internal server error"
		responseData.Err = err

		helper.WriteResponse(w, responseData)
		return
	}

	responseData.StatusCode = 200
	responseData.Message = "todos found"
	responseData.Data = todoResponses

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoIdString := params.ByName("todoId")

	todoId, errCastToInt := strconv.Atoi(todoIdString)

	responseData := helper.ResponseData{}

	if errCastToInt != nil {
		responseData.StatusCode = 500
		responseData.Message = "failed to cast todo id to int"
		responseData.Err = errCastToInt

		helper.WriteResponse(w, responseData)
		return
	}

	todoUpdateRequest := request.TodoUpdateRequest{
		Id: todoId,
	}

	if errReadBody := helper.ReadRequestBody(r, &todoUpdateRequest); errReadBody != nil {
		responseData.StatusCode = 500
		responseData.Message = "failed to read request body"
		responseData.Err = errReadBody

		helper.WriteResponse(w, responseData)
		return
	}

	err := todoController.todoService.Update(r.Context(), todoUpdateRequest)

	if err != nil {
		if errValidation, ok := err.(validator.ValidationErrors); ok {
			responseData.StatusCode = 400
			responseData.Message = "validation error"
			responseData.Data = errValidation.Error()
		} else if errors.Is(helper.ErrRowsNotAffected, err) {
			responseData.StatusCode = 500
			responseData.Message = "no rows affected"
			responseData.Err = err
		} else {
			responseData.StatusCode = 500
			responseData.Message = "internal server error"
			responseData.Err = err
		}

		helper.WriteResponse(w, responseData)
		return
	}

	responseData.StatusCode = 204

	helper.WriteResponse(w, responseData)
}

func (todoController *TodoControllerImpl) Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoIdString := params.ByName("todoId")

	todoId, errCastToInt := strconv.Atoi(todoIdString)

	responseData := helper.ResponseData{}

	if errCastToInt != nil {
		responseData.StatusCode = 500
		responseData.Message = "failed to cast todo id to int"
		responseData.Err = errCastToInt

		helper.WriteResponse(w, responseData)
		return
	}

	err := todoController.todoService.Remove(r.Context(), todoId)

	if err != nil {
		if errors.Is(helper.ErrRowsNotAffected, err) {
			responseData.StatusCode = 500
			responseData.Message = "no rows affected"
			responseData.Err = err
		} else {
			responseData.StatusCode = 500
			responseData.Message = "internal server error"
			responseData.Err = err
		}

		helper.WriteResponse(w, responseData)
		return
	}

	responseData.StatusCode = 204

	helper.WriteResponse(w, responseData)
}
