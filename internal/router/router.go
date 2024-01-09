package router

import (
	"go_todo_api/internal/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, todoController controller.TodoController, authController controller.AuthController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/login", authController.Login)

	router.POST("/api/user", userController.CreateUser)
	router.GET("/api/user/:userId", userController.Get)
	router.GET("/api/users", userController.GetAll)
	router.PUT("/api/user/:userId", userController.Update)
	router.DELETE("/api/user/:userId", userController.Remove)

	router.POST("/api/todo", todoController.CreateTodo)

	router.GET("/api/user/:userId/todo", todoController.GetUserTodos)

	router.GET("/api/todo/:todoId", todoController.Get)
	router.GET("/api/todos", todoController.GetAll)
	router.PUT("/api/todo/:todoId", todoController.Update)
	router.DELETE("/api/todo/:todoId", todoController.Remove)

	return router
}
