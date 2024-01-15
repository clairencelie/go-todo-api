package main

import (
	"database/sql"
	"fmt"
	"go_todo_api/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "192.168.1.9:8080",
		Handler: handler,
	}
}

func NewDB() (*sql.DB, func()) {
	db, _ := database.NewDB(".", false)

	cleanup := func() {
		db.Close()
	}

	return db, cleanup
}

func main() {
	server, closeDb := InitializeServer()

	defer closeDb()

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
