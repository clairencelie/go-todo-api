package controller

import (
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type AuthControllerImpl struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (authController *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userLoginRequest := request.UserLoginRequest{}

	errReadRequestBody := helper.ReadRequestBody(r, &userLoginRequest)

	if errReadRequestBody != nil {
		helper.WriteErrorResponse(w, errReadRequestBody)
		return
	}

	userResponse, err := authController.authService.Login(r.Context(), userLoginRequest)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "login success",
		Data:       userResponse,
	}

	helper.WriteResponse(w, responseData)
}
