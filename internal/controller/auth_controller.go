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
	RefreshToken(w http.ResponseWriter, r *http.Request, params httprouter.Params)
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

	loginResponse, err := authController.authService.Login(r.Context(), userLoginRequest)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "login success",
		Data:       loginResponse,
	}

	helper.WriteResponse(w, responseData)
}

func (authController *AuthControllerImpl) RefreshToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	refreshTokenRequest := request.RefreshTokenRequest{}

	errReadRequestBody := helper.ReadRequestBody(r, &refreshTokenRequest)

	if errReadRequestBody != nil {
		helper.WriteErrorResponse(w, errReadRequestBody)
		return
	}

	refreshTokenResponse, err := authController.authService.RefreshToken(r.Context(), refreshTokenRequest)

	if err != nil {
		helper.WriteErrorResponse(w, err)
		return
	}

	responseData := helper.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "refresh token success",
		Data:       refreshTokenResponse,
	}

	helper.WriteResponse(w, responseData)
}
