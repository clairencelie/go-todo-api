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

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (userController *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := request.UserCreateRequest{}

	if errReadBody := helper.ReadRequestBody(r, &userCreateRequest); errReadBody != nil {
		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "failed to read request body",
			Err:        errReadBody,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	if errCreateUser := userController.userService.Create(r.Context(), userCreateRequest); errCreateUser != nil {
		if validationErrors, ok := errCreateUser.(validator.ValidationErrors); ok {
			responseData := helper.ResponseData{
				StatusCode: 400,
				Message:    "validation error",
				Err:        validationErrors,
			}
			helper.WriteResponse(w, responseData)
			return
		}

		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "failed to create new user",
			Err:        errCreateUser,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: 201,
		Message:    "new user created",
	}

	helper.WriteResponse(w, responseData)
}

func (userController *UserControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIdString := params.ByName("userId")

	userId, errCastToInt := strconv.Atoi(userIdString)

	if errCastToInt != nil {
		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "failed to cast user id to int",
			Err:        errCastToInt,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	userResponse, err := userController.userService.Find(r.Context(), userId)

	if err != nil {
		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "internal server error",
			Err:        err,
		}

		if errors.Is(repository.ErrNotFound, err) {
			responseData.StatusCode = 404
			responseData.Message = "user not found"
		}

		helper.WriteResponse(w, responseData)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: 200,
		Message:    "user found",
		Data:       userResponse,
	}
	helper.WriteResponse(w, responseData)
}

func (userController *UserControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userResponses, err := userController.userService.FindAll(r.Context())

	if err != nil {
		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "internal server error",
			Err:        err,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: 200,
		Message:    "users found",
		Data:       userResponses,
	}
	helper.WriteResponse(w, responseData)
}

func (userController *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIdString := params.ByName("userId")

	userId, errCastToInt := strconv.Atoi(userIdString)

	if errCastToInt != nil {
		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "failed to cast used id",
			Err:        errCastToInt,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	userUpdateRequest := request.UserUpdateRequest{
		Id: userId,
	}

	if errReadBody := helper.ReadRequestBody(r, &userUpdateRequest); errReadBody != nil {
		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "failed to read request body",
			Err:        errReadBody,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	err := userController.userService.Update(r.Context(), userUpdateRequest)

	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			responseData := helper.ResponseData{
				StatusCode: 400,
				Message:    "validation error",
				Err:        validationErrors,
			}
			helper.WriteResponse(w, responseData)
			return
		}

		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "internal server error",
			Err:        err,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: 204,
	}
	helper.WriteResponse(w, responseData)
}

func (userController *UserControllerImpl) Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIdString := params.ByName("userId")

	userId, errCastToInt := strconv.Atoi(userIdString)

	if errCastToInt != nil {
		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "failed to cast used id",
			Err:        errCastToInt,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	err := userController.userService.Remove(r.Context(), userId)

	if err != nil {
		responseData := helper.ResponseData{
			StatusCode: 500,
			Message:    "internal server error",
			Err:        err,
		}
		helper.WriteResponse(w, responseData)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: 204,
	}
	helper.WriteResponse(w, responseData)
}
