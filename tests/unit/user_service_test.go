package unit

import (
	"context"
	"database/sql"
	"go_todo_api/internal/model/entity"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/service"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (mock *UserRepositoryMock) Get(ctx context.Context, db *sql.DB, userId int) (entity.User, error) {
	args := mock.Called(ctx, db, userId)

	if args.Get(1) != nil {
		return args.Get(0).(entity.User), args.Get(1).(error)
	}

	return args.Get(0).(entity.User), nil
}

func (mock *UserRepositoryMock) GetByUsername(ctx context.Context, db *sql.DB, userName string) (entity.User, error) {
	args := mock.Called(ctx, db, userName)

	if args.Get(1) != nil {
		return args.Get(0).(entity.User), args.Get(1).(error)
	}

	return args.Get(0).(entity.User), nil
}

func (mock *UserRepositoryMock) Insert(ctx context.Context, db *sql.DB, user request.UserCreateRequest) error {
	args := mock.Called(ctx, db, user)
	return args.Error(0)
}

func (mock *UserRepositoryMock) Update(ctx context.Context, db *sql.DB, user request.UserUpdateRequest) error {
	args := mock.Called(ctx, db, user)
	return args.Error(0)
}

func (mock *UserRepositoryMock) Delete(ctx context.Context, tx *sql.Tx, userId int) error {
	args := mock.Called(ctx, tx, userId)
	return args.Error(0)
}

func (mock *UserRepositoryMock) DeleteUserTodo(ctx context.Context, tx *sql.Tx, userId int) error {
	args := mock.Called(ctx, tx, userId)
	return args.Error(0)
}

type ValidatorMock struct {
	mock.Mock
}

func (mock *ValidatorMock) StructCtx(ctx context.Context, s interface{}) error {
	args := mock.Called(ctx, s)
	return args.Error(0)
}

func hashPasswordMock(password string) (string, error) {
	return password, nil
}

func TestUserServiceFind(t *testing.T) {
	db, _, errSqlMock := sqlmock.New()

	assert.NoError(t, errSqlMock)

	defer db.Close()

	userRepositoryMock := new(UserRepositoryMock)
	validatorMock := new(ValidatorMock)

	userService := service.NewUserService(db, userRepositoryMock, validatorMock, hashPasswordMock)

	expectedUser := entity.User{
		Id:          1,
		Username:    "apollo",
		Password:    "secret",
		Name:        "Apollo",
		Email:       "apollo@example.xyz",
		PhoneNumber: "081746219124",
		CreatedAt:   "2020-10-10 10:10:10",
		UpdatedAt:   "2020-10-10 10:10:10",
	}

	ctx := context.Background()
	userRepositoryMock.On("Get", ctx, db, 1).Return(expectedUser, nil)

	userResponse, err := userService.Find(ctx, 1)

	assert.NoError(t, err)

	assert.Equal(t, expectedUser.Id, userResponse.Id)
}

func TestUserServiceCreate(t *testing.T) {
	db, _, errSqlMock := sqlmock.New()

	assert.NoError(t, errSqlMock)

	defer db.Close()

	userRepositoryMock := new(UserRepositoryMock)
	validatorMock := new(ValidatorMock)

	userService := service.NewUserService(db, userRepositoryMock, validatorMock, hashPasswordMock)

	userCreateRequest := request.UserCreateRequest{
		Username:    "anto",
		Password:    "secret",
		Name:        "Antonius",
		Email:       "anto@example.xyz",
		PhoneNumber: "08582198125",
	}

	ctx := context.Background()
	validatorMock.On("StructCtx", ctx, userCreateRequest).Return(nil)
	userRepositoryMock.On("Insert", ctx, db, userCreateRequest).Return(nil)

	err := userService.Create(ctx, userCreateRequest)

	assert.NoError(t, err)
}

func TestUserServiceUpdate(t *testing.T) {
	db, _, errSqlMock := sqlmock.New()

	assert.NoError(t, errSqlMock)

	defer db.Close()

	userRepositoryMock := new(UserRepositoryMock)
	validatorMock := new(ValidatorMock)

	userService := service.NewUserService(db, userRepositoryMock, validatorMock, hashPasswordMock)

	userUpdateRequest := request.UserUpdateRequest{
		Id:          1,
		Username:    "budiman",
		Name:        "Budiman",
		Email:       "budiman@example.xyz",
		PhoneNumber: "0123456789",
	}

	ctx := context.Background()
	validatorMock.On("StructCtx", ctx, userUpdateRequest).Return(nil)
	userRepositoryMock.On("Update", ctx, db, userUpdateRequest).Return(nil)

	err := userService.Update(ctx, userUpdateRequest)

	assert.NoError(t, err)
}

func TestUserServiceDelete(t *testing.T) {
	db, mockDB, errSqlMock := sqlmock.New()

	assert.NoError(t, errSqlMock)

	defer db.Close()

	mockDB.ExpectBegin()
	mockDB.ExpectCommit()

	userRepositoryMock := new(UserRepositoryMock)

	ctx := context.Background()
	userRepositoryMock.On("DeleteUserTodo", ctx, mock.AnythingOfType("*sql.Tx"), 1).Return(nil)
	userRepositoryMock.On("Delete", ctx, mock.AnythingOfType("*sql.Tx"), 1).Return(nil)

	validatorMock := new(ValidatorMock)

	userService := service.NewUserService(db, userRepositoryMock, validatorMock, hashPasswordMock)

	err := userService.Remove(ctx, 1)
	assert.NoError(t, err)

	errMock := mockDB.ExpectationsWereMet()
	assert.NoError(t, errMock)
}
