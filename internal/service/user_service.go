package service

import (
	"context"
	"database/sql"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/model/response"
	"go_todo_api/internal/repository"
	customvalidator "go_todo_api/internal/validator"
)

type UserService interface {
	Find(ctx context.Context, userId int) (response.UserResponse, error)
	Create(ctx context.Context, user request.UserCreateRequest) error
	Update(ctx context.Context, user request.UserUpdateRequest) error
	Remove(ctx context.Context, userId int) error
}

type UserServiceImpl struct {
	db             *sql.DB
	userRepository repository.UserRepository
	validate       customvalidator.CustomValidator
	passwordHasher func(password string) (string, error)
}

func NewUserService(db *sql.DB, userRepository repository.UserRepository, validate customvalidator.CustomValidator, passwordHasher func(password string) (string, error)) UserService {
	return &UserServiceImpl{
		db:             db,
		userRepository: userRepository,
		validate:       validate,
		passwordHasher: passwordHasher,
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

func (userService *UserServiceImpl) Create(ctx context.Context, user request.UserCreateRequest) error {
	if err := userService.validate.StructCtx(ctx, user); err != nil {
		return err
	}

	hashedPassword, errHashingPassword := userService.passwordHasher(user.Password)

	if errHashingPassword != nil {
		return errHashingPassword
	}

	user.Password = hashedPassword

	err := userService.userRepository.Insert(ctx, userService.db, user)

	if err != nil {
		return err
	}

	return nil
}

func (userService *UserServiceImpl) Update(ctx context.Context, user request.UserUpdateRequest) error {
	if err := userService.validate.StructCtx(ctx, user); err != nil {
		return err
	}

	err := userService.userRepository.Update(ctx, userService.db, user)

	if err != nil {
		return err
	}

	return nil
}

func (userService *UserServiceImpl) Remove(ctx context.Context, userId int) error {
	tx, errTxBegin := userService.db.Begin()

	if errTxBegin != nil {
		return errTxBegin
	}

	errTodoDelete := userService.userRepository.DeleteUserTodo(ctx, tx, userId)
	if errTodoDelete != nil {
		tx.Rollback()
		return errTodoDelete
	}

	err := userService.userRepository.Delete(ctx, tx, userId)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
