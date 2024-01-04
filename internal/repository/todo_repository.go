package repository

import (
	"context"
	"database/sql"
	"go_todo_api/internal/model/entity"
	"go_todo_api/internal/model/request"
)

type TodoRepository interface {
	Get(ctx context.Context, db *sql.DB, todoId int) entity.Todo
	GetAll(ctx context.Context) []entity.Todo
	Insert(ctx context.Context, tx *sql.Tx, todo request.TodoCreateRequest) error
	Update(ctx context.Context, tx *sql.Tx, todo request.TodoUpdateRequest) error
	Delete(ctx context.Context, tx *sql.Tx) error
}
