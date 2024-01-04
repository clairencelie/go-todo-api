package controller

import (
	"go_todo_api/internal/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (userController *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (userController *UserControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (userController *UserControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (userController *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (userController *UserControllerImpl) Remove(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
