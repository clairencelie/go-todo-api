package repository

import (
	"context"
	"database/sql"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/model/entity"
	"go_todo_api/internal/model/request"
)

type TodoRepository interface {
	Get(ctx context.Context, db *sql.DB, todoId int) (entity.Todo, error)
	GetAll(ctx context.Context, db *sql.DB) ([]entity.Todo, error)
	Insert(ctx context.Context, tx *sql.Tx, todo request.TodoCreateRequest) error
	Update(ctx context.Context, tx *sql.Tx, todo request.TodoUpdateRequest) error
	Delete(ctx context.Context, tx *sql.Tx, todoId int) error
}

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (repository TodoRepositoryImpl) Get(ctx context.Context, db *sql.DB, todoId int) (entity.Todo, error) {
	query := "SELECT id, user_id, title, description, is_done, created_at, updated_at FROM todos WHERE id = ? LIMIT 1"

	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		return entity.Todo{}, err
	}

	rows, queryErr := stmt.QueryContext(ctx, todoId)

	if queryErr != nil {
		return entity.Todo{}, queryErr
	}

	defer rows.Close()

	if rows.Next() {
		todo := entity.Todo{}

		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Description, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt)

		if err != nil {
			return entity.Todo{}, err
		}

		return todo, nil
	}

	return entity.Todo{}, ErrNotFound
}

func (repository TodoRepositoryImpl) GetAll(ctx context.Context, db *sql.DB) ([]entity.Todo, error) {
	query := "SELECT id, user_id, title, description, is_done, created_at, updated_at FROM todos"

	rows, queryErr := db.Query(query)

	if queryErr != nil {
		return nil, queryErr
	}

	defer rows.Close()

	todos := []entity.Todo{}

	for rows.Next() {
		todo := entity.Todo{}

		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Description, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (repository TodoRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, todo request.TodoCreateRequest) error {
	query := "INSERT INTO todos (user_id, title, description) VALUES (?, ?, ?)"

	stmt, errPrepare := tx.PrepareContext(ctx, query)

	if errPrepare != nil {
		return errPrepare
	}

	sqlResult, errExec := stmt.ExecContext(ctx, todo.UserId, todo.Title, todo.Description)

	if errExec != nil {
		return errExec
	}

	err := helper.CheckRowsAffected(sqlResult)

	if err != nil {
		return err
	}

	return nil
}

func (repository TodoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo request.TodoUpdateRequest) error {
	query := "UPDATE todos SET title=?, description=?, is_done=? WHERE id=?"

	stmt, errPrepare := tx.PrepareContext(ctx, query)

	if errPrepare != nil {
		return errPrepare
	}

	sqlResult, errExec := stmt.ExecContext(ctx, todo.Title, todo.Description, todo.IsDone, todo.Id)

	if errExec != nil {
		return errExec
	}

	err := helper.CheckRowsAffected(sqlResult)

	if err != nil {
		return err
	}

	return nil
}

func (repository TodoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todoId int) error {
	query := "DELETE FROM todos WHERE id = ?"

	stmt, errPrepare := tx.PrepareContext(ctx, query)

	if errPrepare != nil {
		return errPrepare
	}

	sqlResult, errExec := stmt.ExecContext(ctx, todoId)

	if errExec != nil {
		return errExec
	}

	err := helper.CheckRowsAffected(sqlResult)

	if err != nil {
		return err
	}

	return nil
}
