package router

import (
	"go_todo_api/internal/controller"
	"go_todo_api/internal/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, todoController controller.TodoController, authController controller.AuthController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/login", authController.Login)
	router.POST("/api/token/refresh", authController.RefreshToken)

	router.POST("/api/user", middleware.AuthMiddleware(userController.CreateUser))
	router.GET("/api/user/:userId", middleware.AuthMiddleware(userController.Get))
	router.PUT("/api/user/:userId", middleware.AuthMiddleware(userController.Update))
	router.DELETE("/api/user/:userId", middleware.AuthMiddleware(userController.Remove))

	router.POST("/api/todo", middleware.AuthMiddleware(todoController.CreateTodo))
	router.GET("/api/user/:userId/todo", middleware.AuthMiddleware(todoController.GetUserTodos))
	router.GET("/api/todo/:todoId", middleware.AuthMiddleware(todoController.Get))
	router.PUT("/api/todo/:todoId", middleware.AuthMiddleware(todoController.Update))
	router.PATCH("/api/todo/completion/:todoId", middleware.AuthMiddleware(todoController.UpdateTodoCompletion))
	router.DELETE("/api/todo/:todoId", middleware.AuthMiddleware(todoController.Remove))

	return router
}
