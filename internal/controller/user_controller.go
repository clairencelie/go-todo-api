package controller

import (
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
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
		helper.WriteErrorResponse(w, errReadBody)
		return
	}

	errCreateUser := userController.userService.Create(r.Context(), userCreateRequest)

	if errCreateUser != nil {
		helper.WriteErrorResponse(w, errCreateUser)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    "new user created",
	}

	helper.WriteResponse(w, responseData)
}

func (userController *UserControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIdString := params.ByName("userId")

	userId, errCastToInt := strconv.Atoi(userIdString)

	if errCastToInt != nil {
		helper.WriteErrorResponse(w, errCastToInt)
		return
	}

	userResponse, err := userController.userService.Find(r.Context(), userId)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "user found",
		Data:       userResponse,
	}

	helper.WriteResponse(w, responseData)
}

func (userController *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIdString := params.ByName("userId")

	userId, errCastToInt := strconv.Atoi(userIdString)

	if errCastToInt != nil {
		helper.WriteErrorResponse(w, errCastToInt)
		return
	}

	userUpdateRequest := request.UserUpdateRequest{
		Id: userId,
	}

	if errReadBody := helper.ReadRequestBody(r, &userUpdateRequest); errReadBody != nil {
		helper.WriteErrorResponse(w, errReadBody)
		return
	}

	err := userController.userService.Update(r.Context(), userUpdateRequest)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusNoContent,
	}

	helper.WriteResponse(w, responseData)
}

func (userController *UserControllerImpl) Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIdString := params.ByName("userId")

	userId, errCastToInt := strconv.Atoi(userIdString)

	if errCastToInt != nil {
		helper.WriteErrorResponse(w, errCastToInt)
		return
	}

	err := userController.userService.Remove(r.Context(), userId)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusNoContent,
	}

	helper.WriteResponse(w, responseData)
}
