package main

import (
	"fmt"
	"go_todo_api/database"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/router"
	"go_todo_api/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := database.NewDB(".")

	validate := validator.New()

	userService := service.NewUserService(db, repository.NewUserRepository(), validate)
	userController := controller.NewUserController(userService)

	todoService := service.NewTodoService(db, repository.NewTodoRepository(), validate)
	todoController := controller.NewTodoController(todoService)

	router := router.NewRouter(userController, todoController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
