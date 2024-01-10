package integration

import (
	"database/sql"
	"go_todo_api/database"
	"strconv"
)

var TestDb, _ = database.NewDB("./../../", true)

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
