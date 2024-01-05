package router

import (
	"go_todo_api/internal/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/user", userController.CreateUser)
	router.GET("/api/user/:userId", userController.Get)
	router.GET("/api/users", userController.GetAll)
	router.PUT("/api/user/:userId", userController.Update)
	router.DELETE("/api/user/:userId", userController.Remove)

	return router
}
