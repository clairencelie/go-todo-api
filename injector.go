//go:build wireinject
// +build wireinject

package main

import (
	"go_todo_api/internal/controller"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/middleware"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/router"
	"go_todo_api/internal/service"
	"go_todo_api/internal/validator"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	helper.HashFunction,
	service.NewUserService,
	controller.NewUserController,
)

var authSet = wire.NewSet(
	service.NewAuthService,
	controller.NewAuthController,
)

var todoSet = wire.NewSet(
	repository.NewTodoRepository,
	service.NewTodoService,
	controller.NewTodoController,
)

func InitializeServer() (*http.Server, func()) {
	wire.Build(
		NewDB,
		validator.NewValidator,
		userSet,
		authSet,
		todoSet,
		router.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewLogMiddleware,
		NewServer,
	)

	return nil, nil
}
