package repository

import (
	"context"
	"database/sql"
	"go_todo_api/internal/model/entity"
	"go_todo_api/internal/model/request"
)

type UserRepository interface {
	Get(ctx context.Context, userId int) entity.User
	GetAll(ctx context.Context) []entity.User
	Insert(ctx context.Context, tx *sql.Tx, user request.UserCreateRequest) error
	Update(ctx context.Context, tx *sql.Tx, user request.UserUpdateRequest) error
	Delete(ctx context.Context, tx *sql.Tx) error
}
