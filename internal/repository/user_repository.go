package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_todo_api/internal/model/entity"
	"go_todo_api/internal/model/request"
)

type UserRepository interface {
	Get(ctx context.Context, db *sql.DB, userId int) (entity.User, error)
	GetAll(ctx context.Context, db *sql.DB) ([]entity.User, error)
	Insert(ctx context.Context, tx *sql.Tx, user request.UserCreateRequest) error
	Update(ctx context.Context, tx *sql.Tx, user request.UserUpdateRequest) error
	Delete(ctx context.Context, tx *sql.Tx) error
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

var (
	errUserNotFound = errors.New("user not found")
)

func (repository UserRepositoryImpl) Get(ctx context.Context, db *sql.DB, userId int) (entity.User, error) {
	query := "SELECT id, username, password, name, email, phone_number, created_at, updated_at FROM users WHERE id = ? LIMIT 1"

	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		return entity.User{}, err
	}

	rows, queryErr := stmt.QueryContext(ctx, userId)

	if queryErr != nil {
		return entity.User{}, queryErr
	}

	defer rows.Close()

	if rows.Next() {
		user := entity.User{}

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.Email, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return entity.User{}, err
		}

		return user, nil
	}

	return entity.User{}, errUserNotFound
}

func (repository UserRepositoryImpl) GetAll(ctx context.Context, db *sql.DB) ([]entity.User, error) {
	return []entity.User{}, nil
}

func (repository UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user request.UserCreateRequest) error {
	return nil
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user request.UserUpdateRequest) error {
	return nil
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx) error {
	return nil
}
