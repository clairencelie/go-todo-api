package service

import (
	"context"
	"database/sql"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/model/response"
	"go_todo_api/internal/repository"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Find(ctx context.Context, userId int) (response.UserResponse, error)
	FindAll(ctx context.Context) ([]response.UserResponse, error)
	Create(ctx context.Context, user request.UserCreateRequest) error
	Update(ctx context.Context, user request.UserUpdateRequest) error
	Remove(ctx context.Context, userId int) error
}

type UserServiceImpl struct {
	db             *sql.DB
	userRepository repository.UserRepository
	validate       *validator.Validate
}

func NewUserService(db *sql.DB, userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		db:             db,
		userRepository: userRepository,
		validate:       validate,
	}
}

func (userService *UserServiceImpl) Find(ctx context.Context, userId int) (response.UserResponse, error) {
	user, err := userService.userRepository.Get(ctx, userService.db, userId)

	if err != nil {
		return response.UserResponse{}, err
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

func (userService *UserServiceImpl) FindAll(ctx context.Context) ([]response.UserResponse, error) {
	users, err := userService.userRepository.GetAll(ctx, userService.db)

	if err != nil {
		return nil, err
	}

	userResponses := []response.UserResponse{}

	for _, user := range users {
		userResponse := response.UserResponse{
			Id:          user.Id,
			Username:    user.Username,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			CreatedAt:   user.CreatedAt,
		}

		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (userService *UserServiceImpl) Create(ctx context.Context, user request.UserCreateRequest) error {
	if err := userService.validate.StructCtx(ctx, user); err != nil {
		return err
	}

	tx, errTxBegin := userService.db.Begin()

	if errTxBegin != nil {
		return errTxBegin
	}

	// need to validate user input later
	err := userService.userRepository.Insert(ctx, tx, user)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (userService *UserServiceImpl) Update(ctx context.Context, user request.UserUpdateRequest) error {
	tx, errTxBegin := userService.db.Begin()

	if errTxBegin != nil {
		return errTxBegin
	}

	// need to validate user input later
	err := userService.userRepository.Update(ctx, tx, user)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (userService *UserServiceImpl) Remove(ctx context.Context, userId int) error {
	tx, errTxBegin := userService.db.Begin()

	if errTxBegin != nil {
		return errTxBegin
	}

	// need to validate user input later
	err := userService.userRepository.Delete(ctx, tx, userId)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
