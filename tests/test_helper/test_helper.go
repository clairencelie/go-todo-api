package testhelper

import (
	"database/sql"
	"go_todo_api/internal/helper"
	"strconv"
)

func ResetDB(testDb *sql.DB) {
	testDb.Exec("DELETE FROM todos")
	testDb.Exec("DELETE FROM users")
}

func InsertSingleUser(testDb *sql.DB) int64 {
	hashedPassword, errHashingPassword := helper.HashPassword("rahasia")

	if errHashingPassword != nil {
		panic(errHashingPassword)
	}

	userSqlResult, errExecUser := testDb.Exec("INSERT INTO users (username, password, name, email, phone_number) VALUES ('budi', ?, 'Budi', 'budi@example.xyz', '081234567')", hashedPassword)

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
