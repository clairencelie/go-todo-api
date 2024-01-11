package integration

import (
	"context"
	"database/sql"
	"go_todo_api/database"
	"go_todo_api/internal/controller"
	"go_todo_api/internal/repository"
	"go_todo_api/internal/router"
	"go_todo_api/internal/service"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var TestDb, _ = database.NewDB("./../../", true)
var Ctx = context.Background()

var TodoRepository = repository.NewTodoRepository()
var UserRepository = repository.NewUserRepository()

var Validator = validator.New()
var UserService = service.NewUserService(TestDb, UserRepository, Validator)
var AuthService = service.NewAuthService(TestDb, UserRepository, Validator)
var TodoService = service.NewTodoService(TestDb, TodoRepository, Validator)

var UserController = controller.NewUserController(UserService)
var AuthController = controller.NewAuthController(AuthService)
var TodoController = controller.NewTodoController(TodoService)

var Router = router.NewRouter(UserController, TodoController, AuthController)

func ResetDB() {
	TestDb.Exec("DELETE FROM todos")
	TestDb.Exec("DELETE FROM users")
}

func InsertSingleUser(testDb *sql.DB) int64 {
	userSqlResult, errExecUser := testDb.Exec("INSERT INTO users (username, password, name, email, phone_number) VALUES ('budi', 'rahasia', 'Budi', 'budi@example.xyz', '081234567')")

	if errExecUser != nil {
		panic(errExecUser)
	}

	userLastInsertId, errUserLastInsertId := userSqlResult.LastInsertId()

	if errUserLastInsertId != nil {
		panic(errUserLastInsertId)
	}

	return userLastInsertId
}

func InsertManyUser(testDb *sql.DB, count int) {
	for i := 1; i <= count; i++ {
		_, errExecUser := testDb.Exec("INSERT INTO users (username, password, name, email, phone_number) VALUES (?, 'rahasia', ?, ?, ?)", "budi"+strconv.Itoa(i), "Budi "+strconv.Itoa(i), "budi"+strconv.Itoa(i)+"@example.xyz", "0812345"+strconv.Itoa(i))

		if errExecUser != nil {
			panic(errExecUser)
		}
	}
}

func InsertSingleTodo(testDb *sql.DB) int64 {
	userSqlResult, errExecUser := testDb.Exec("INSERT INTO users (username, password, name, email, phone_number) VALUES ('budi', 'rahasia', 'Budi', 'budi@example.xyz', '081234567')")

	if errExecUser != nil {
		panic(errExecUser)
	}

	userLastInsertId, errUserLastInsertId := userSqlResult.LastInsertId()

	if errUserLastInsertId != nil {
		panic(errUserLastInsertId)
	}

	todoSqlResult, errExecTodo := testDb.Exec("INSERT INTO todos (user_id, title, description) VALUES (?, ?, ?)", userLastInsertId, "todo 1", "deskripsi todo 1")

	if errExecTodo != nil {
		panic(errExecTodo)
	}

	todoLastInsertId, errTodoLastInsertId := todoSqlResult.LastInsertId()

	if errTodoLastInsertId != nil {
		panic(errTodoLastInsertId)
	}

	return todoLastInsertId
}

func InsertManyTodo(testDb *sql.DB, count int) {
	userSqlResult, errExecUser := testDb.Exec("INSERT INTO users (username, password, name, email, phone_number) VALUES ('budi', 'rahasia', 'Budi', 'budi@example.xyz', '081234567')")

	if errExecUser != nil {
		panic(errExecUser)
	}

	userLastInsertId, errUserLastInsertId := userSqlResult.LastInsertId()

	if errUserLastInsertId != nil {
		panic(errUserLastInsertId)
	}

	for i := 1; i <= count; i++ {
		_, errExecTodo := testDb.Exec("INSERT INTO todos (user_id, title, description) VALUES (?, ?, ?)", userLastInsertId, "todo "+strconv.Itoa(i), "deskripsi todo "+strconv.Itoa(i))

		if errExecTodo != nil {
			panic(errExecTodo)
		}
	}
}
