package service

import (
	"context"
	"database/sql"
	"errors"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/model/response"
	"go_todo_api/internal/repository"

	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Login(ctx context.Context, loginRequest request.UserLoginRequest) (response.UserResponse, error)
}

type AuthServiceImpl struct {
	db             *sql.DB
	userRepository repository.UserRepository
	validate       *validator.Validate
}

func NewAuthService(db *sql.DB, userRepository repository.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		db:             db,
		userRepository: userRepository,
		validate:       validate,
	}
}

func (authService *AuthServiceImpl) Login(ctx context.Context, loginRequest request.UserLoginRequest) (response.UserResponse, error) {
	errValidation := authService.validate.StructCtx(ctx, loginRequest)

	if errValidation != nil {
		return response.UserResponse{}, errValidation
	}

	user, err := authService.userRepository.GetByUsername(ctx, authService.db, loginRequest.Username)

	if err != nil {
		if errors.Is(helper.ErrNotFound, err) {
			return response.UserResponse{}, helper.ErrLoginFailed
		}
		return response.UserResponse{}, err
	}

	if !helper.CheckPassword(loginRequest.Password, user.Password) {
		return response.UserResponse{}, helper.ErrLoginFailed
	}

	userResponse := response.UserResponse{
		Id:          user.Id,
		Username:    user.Username,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}

	return userResponse, nil
}
