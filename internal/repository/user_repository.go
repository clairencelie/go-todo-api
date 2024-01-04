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
	Delete(ctx context.Context, tx *sql.Tx, userId int) error
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

var (
	errUserNotFound    = errors.New("user not found")
	errRowsNotAffected = errors.New("no rows affected")
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
	query := "SELECT id, username, password, name, email, phone_number, created_at, updated_at FROM users"

	rows, queryErr := db.Query(query)

	if queryErr != nil {
		return nil, queryErr
	}

	defer rows.Close()

	users := []entity.User{}

	for rows.Next() {
		user := entity.User{}

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.Email, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user request.UserCreateRequest) error {
	query := "INSERT INTO users (username, password, name, email, phone_number) VALUES (?, ?, ?, ?, ?)"

	stmt, errPrepare := tx.PrepareContext(ctx, query)

	if errPrepare != nil {
		return errPrepare
	}

	sqlResult, errExec := stmt.ExecContext(ctx, user.Username, user.Password, user.Name, user.Email, user.PhoneNumber)

	if errExec != nil {
		return errExec
	}

	err := checkRowsAffected(sqlResult)

	if err != nil {
		return err
	}

	return nil
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user request.UserUpdateRequest) error {
	query := "UPDATE users SET username=?, name=?, email=?, phone_number=? WHERE id=?"

	stmt, errPrepare := tx.PrepareContext(ctx, query)

	if errPrepare != nil {
		return errPrepare
	}

	sqlResult, errExec := stmt.ExecContext(ctx, user.Username, user.Name, user.Email, user.PhoneNumber, user.Id)

	if errExec != nil {
		return errExec
	}

	err := checkRowsAffected(sqlResult)

	if err != nil {
		return err
	}

	return nil
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) error {
	query := "DELETE FROM users WHERE id = ?"

	stmt, errPrepare := tx.PrepareContext(ctx, query)

	if errPrepare != nil {
		return errPrepare
	}

	sqlResult, errExec := stmt.ExecContext(ctx, userId)

	if errExec != nil {
		return errExec
	}

	err := checkRowsAffected(sqlResult)

	if err != nil {
		return err
	}

	return nil
}

func checkRowsAffected(sqlResult sql.Result) error {
	rowsAffected, err := sqlResult.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errRowsNotAffected
	}

	return nil
}
