package main

import (
	"context"
	"database/sql"
	"fmt"
	"go_todo_api/database"
	"net/http"
	"os"
	"os/signal"

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

	defer func() {
		fmt.Println("Closing DB...")
		closeDb()
		fmt.Println("DB Closed...")
		cancel()
	}()

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err.Error())
			cancel()
		}
	}()

	<-ctx.Done()

	fmt.Println("Cleaning App")
}
