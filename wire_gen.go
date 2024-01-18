// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/helper"
	"go_todo_api/internal/middleware"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/router"
	"go_todo_api/internal/service"
	"go_todo_api/internal/validator"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializeServer() (*http.Server, func()) {
	db, cleanup := NewDB()
	userRepository := repository.NewUserRepository()
	customValidator := validator.NewValidator()
	v := helper.HashFunction()
	userService := service.NewUserService(db, userRepository, customValidator, v)
	userController := controller.NewUserController(userService)
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(db, todoRepository, customValidator)
	todoController := controller.NewTodoController(todoService)
	authService := service.NewAuthService(db, userRepository, customValidator)
	authController := controller.NewAuthController(authService)
	httprouterRouter := router.NewRouter(userController, todoController, authController)
	logMiddlewareHandler := middleware.NewLogMiddleware(httprouterRouter)
	server := NewServer(logMiddlewareHandler)
	return server, func() {
		cleanup()
	}
}

// injector.go:

var userSet = wire.NewSet(repository.NewUserRepository, helper.HashFunction, service.NewUserService, controller.NewUserController)

var authSet = wire.NewSet(service.NewAuthService, controller.NewAuthController)

var todoSet = wire.NewSet(repository.NewTodoRepository, service.NewTodoService, controller.NewTodoController)
