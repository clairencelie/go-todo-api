package controller

import (
	"errors"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
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
	responseData := helper.ResponseData{}

	userLoginRequest := request.UserLoginRequest{}

	errReadRequestBody := helper.ReadRequestBody(r, &userLoginRequest)

	if errReadRequestBody != nil {
		responseData.StatusCode = 500
		responseData.Message = "error read request body"
		responseData.Err = errReadRequestBody

		helper.WriteResponse(w, responseData)
		return
	}

	userResponse, err := authController.authService.Login(r.Context(), userLoginRequest)

	if err != nil {
		if errors.Is(service.ErrLoginFailed, err) {
			responseData.StatusCode = 401
			responseData.Message = "login failed"
			responseData.Err = err
		} else if errValidation, ok := err.(validator.ValidationErrors); ok {
			responseData.StatusCode = 400
			responseData.Message = "bad request"
			responseData.Err = errValidation
		} else {
			responseData.StatusCode = 500
			responseData.Message = "internal server error"
			responseData.Err = err
		}

		helper.WriteResponse(w, responseData)
		return
	}

	responseData.StatusCode = 200
	responseData.Message = "login success"
	responseData.Data = userResponse

	helper.WriteResponse(w, responseData)
}
