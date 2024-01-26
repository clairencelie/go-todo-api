package main

import (
	"context"
	"database/sql"
	"fmt"
	"go_todo_api/database"
	"go_todo_api/internal/middleware"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewServer(handler middleware.LogMiddlewareHandler) *http.Server {
	errEnvLoad := godotenv.Load("config.env")

	if errEnvLoad != nil {
		fmt.Println(errEnvLoad.Error())
		return nil
	}

	addr := os.Getenv("APP_URL")

	return &http.Server{
		Addr:    addr,
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
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for sig := range c {
			fmt.Println(sig)
			cancel()
		}
	}()

	server, closeDb := InitializeServer()

	go func() {
		fmt.Println("Server running on:", "http://"+server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err.Error())
			cancel()
		}
	}()

	<-ctx.Done()

	fmt.Println("Cleaning App...")
	fmt.Println("Closing DB...")
	closeDb()
	fmt.Println("DB Closed...")
}
