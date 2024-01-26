package service

import (
	"context"
	"database/sql"
	"errors"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/model/response"
	"go_todo_api/internal/repository"
	customvalidator "go_todo_api/internal/validator"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, loginRequest request.UserLoginRequest) (response.LoginResponse, error)
	RefreshToken(ctx context.Context, tokenRefreshRequest request.RefreshTokenRequest) (response.RefreshTokenResponse, error)
}

type AuthServiceImpl struct {
	db             *sql.DB
	userRepository repository.UserRepository
	validate       customvalidator.CustomValidator
}

func NewAuthService(db *sql.DB, userRepository repository.UserRepository, validate customvalidator.CustomValidator) AuthService {
	return &AuthServiceImpl{
		db:             db,
		userRepository: userRepository,
		validate:       validate,
	}
}

func (authService *AuthServiceImpl) Login(ctx context.Context, loginRequest request.UserLoginRequest) (response.LoginResponse, error) {
	errValidation := authService.validate.StructCtx(ctx, loginRequest)

	if errValidation != nil {
		return response.LoginResponse{}, errValidation
	}

	user, err := authService.userRepository.GetByUsername(ctx, authService.db, loginRequest.Username)

	if err != nil {
		if errors.Is(helper.ErrNotFound, err) {
			return response.LoginResponse{}, helper.ErrLoginFailed
		}
		return response.LoginResponse{}, err
	}

	if !helper.CheckPassword(loginRequest.Password, user.Password) {
		return response.LoginResponse{}, helper.ErrLoginFailed
	}

	accessTokenExp := time.Now().Add(time.Duration(15) * time.Minute).Unix()
	refreshTokenExp := time.Now().Add(time.Duration(720) * time.Hour).Unix()

	accessTokenStr, errGenerateAccessToken := helper.GenerateJWT(user.Username, accessTokenExp)

	if errGenerateAccessToken != nil {
		return response.LoginResponse{}, errGenerateAccessToken
	}

	refreshTokenStr, errGenerateRefreshToken := helper.GenerateJWT(user.Username, refreshTokenExp)

	if errGenerateRefreshToken != nil {
		return response.LoginResponse{}, errGenerateRefreshToken
	}

	userResponse := response.UserResponse{
		Id:          user.Id,
		Username:    user.Username,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}

	loginResponse := response.LoginResponse{
		UserResponse: userResponse,
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
	}

	return loginResponse, nil
}

func (authService *AuthServiceImpl) RefreshToken(ctx context.Context, refreshTokenRequest request.RefreshTokenRequest) (response.RefreshTokenResponse, error) {
	errValidation := authService.validate.StructCtx(ctx, refreshTokenRequest)

	if errValidation != nil {
		return response.RefreshTokenResponse{}, errValidation
	}

	// Validate refresh token.
	_, errValidateRefreshToken := helper.ValidateJWT(refreshTokenRequest.RefreshToken)

	if errValidateRefreshToken != nil {
		// Can make new type error: refresh token invalid
		return response.RefreshTokenResponse{}, errValidateRefreshToken
	}

	// Validate access token
	requestAccessToken, errValidateAccessToken := helper.ValidateJWT(refreshTokenRequest.AccessToken)

	refreshTokenResponse := response.RefreshTokenResponse{}

	if errValidateAccessToken != nil {
		if errValidateAccessToken.Error() == "token has invalid claims: token is expired" { // Access token valid, but expired.
			sub, errGetSub := requestAccessToken.Claims.GetSubject()

			if errGetSub != nil {
				return response.RefreshTokenResponse{}, errGetSub
			}

			newAccessTokenExp := time.Now().Add(time.Duration(15) * time.Minute).Unix()
			newAccessTokenStr, errGenerateAccessToken := helper.GenerateJWT(sub, newAccessTokenExp)

			if errGenerateAccessToken != nil {
				return response.RefreshTokenResponse{}, errGenerateAccessToken
			}

			refreshTokenResponse.AccessToken = newAccessTokenStr
		} else { // Invalid access token.
			return response.RefreshTokenResponse{}, errValidateRefreshToken
		}
	}

	return refreshTokenResponse, nil
}
