package unit

import (
	"context"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/request"
	"go_todo_api/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var userRepository = repository.NewUserRepository()

func TestUserRepositoryGetById(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.Nil(t, err)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "name", "email", "phone_number", "created_at", "updated_at"}).
		AddRow(1, "budi", "secret", "Budi", "budi@example.xyz", "087654321", "2024-01-01", "2024-01-01")

	mock.ExpectPrepare("SELECT id, username, password, name, email, phone_number, created_at, updated_at FROM users").ExpectQuery().WithArgs(1).WillReturnRows(rows)

	user, errGetUser := userRepository.Get(context.Background(), db, 1)

	assert.NoError(t, errGetUser)
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "budi", user.Username)

	mock.ExpectPrepare("SELECT id, username, password, name, email, phone_number, created_at, updated_at FROM users").ExpectQuery().WithArgs(2).WillReturnError(helper.ErrNotFound)

	_, errUserNotFound := userRepository.Get(context.Background(), db, 2)

	assert.Error(t, errUserNotFound)
}

func TestUserRepositoryGetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.Nil(t, err)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "name", "email", "phone_number", "created_at", "updated_at"}).
		AddRow(2, "apollo", "secret", "Apollo", "apolo@example.xyz", "09847218", "2024-01-02", "2024-01-02")

	mock.ExpectPrepare("SELECT id, username, password, name, email, phone_number, created_at, updated_at FROM users").ExpectQuery().WithArgs("apollo").WillReturnRows(rows)

	user, errGetByUsername := userRepository.GetByUsername(context.Background(), db, "apollo")

	assert.NoError(t, errGetByUsername)
	assert.Equal(t, 2, user.Id)
	assert.Equal(t, "apollo", user.Username)
	assert.Equal(t, "Apollo", user.Name)

	mock.ExpectPrepare("SELECT id, username, password, name, email, phone_number, created_at, updated_at FROM users").ExpectQuery().WithArgs("unknown_user").WillReturnError(helper.ErrNotFound)

	_, errUserNotFound := userRepository.GetByUsername(context.Background(), db, "unknown_user")

	assert.Error(t, errUserNotFound)
}

func TestUserRepositoryInsert(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.Nil(t, err)

	defer db.Close()

	userCreateRequest := request.UserCreateRequest{
		Username:    "athena",
		Password:    "secret",
		Name:        "Athena",
		Email:       "athena@example.xyz",
		PhoneNumber: "0748592719",
	}

	mock.ExpectPrepare("INSERT INTO users").ExpectExec().WithArgs(userCreateRequest.Username, userCreateRequest.Password, userCreateRequest.Name, userCreateRequest.Email, userCreateRequest.PhoneNumber).WillReturnResult(sqlmock.NewResult(3, 1))

	errUserInsert := userRepository.Insert(context.Background(), db, userCreateRequest)

	assert.NoError(t, errUserInsert)
}

func TestUserRepositoryUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.Nil(t, err)

	defer db.Close()

	userUpdateRequest := request.UserUpdateRequest{
		Id:          1,
		Username:    "budiman",
		Name:        "Budiman",
		Email:       "budiman@example.xyz",
		PhoneNumber: "0123456789",
	}

	mock.ExpectPrepare("UPDATE users").ExpectExec().WithArgs(userUpdateRequest.Username, userUpdateRequest.Name, userUpdateRequest.Email, userUpdateRequest.PhoneNumber, userUpdateRequest.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	errUserUpdate := userRepository.Update(context.Background(), db, userUpdateRequest)

	assert.NoError(t, errUserUpdate)
}

func TestUserRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.Nil(t, err)

	defer db.Close()

	mock.ExpectBegin()

	tx, errTx := db.Begin()

	assert.NoError(t, errTx)

	mock.ExpectPrepare("DELETE FROM users").ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	errUserDelete := userRepository.Delete(context.Background(), tx, 1)

	assert.NoError(t, errUserDelete)
}

func TestUserRepositoryDeleteUserTodo(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.Nil(t, err)

	defer db.Close()

	mock.ExpectBegin()

	tx, errTx := db.Begin()

	assert.NoError(t, errTx)

	mock.ExpectPrepare("DELETE FROM todos").ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	errUserTodoDelete := userRepository.DeleteUserTodo(context.Background(), tx, 1)

	assert.NoError(t, errUserTodoDelete)
}
