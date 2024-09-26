package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/anucha-tk/go_todo/internal/application"
	"github.com/anucha-tk/go_todo/internal/infrastructure/db"
	httphandler "github.com/anucha-tk/go_todo/internal/interfaces/http"
)

func main() {
	port := ":3000"
	flag.StringVar(&port, "port", port, "port to listen on")
	flag.Parse()

	dbConn, err := db.Init()
	if err != nil {
		log.Fatalf("cannot connect database: %v", err)
	}

	todoRepo := db.NewTodoRepositoryGORM(dbConn)
	todoService := application.NewTodoService(todoRepo)
	todoHandler := httphandler.NewTodoHandler(todoService)
	aboutHandler := httphandler.NewAboutHandler()

	router := routes(todoHandler, aboutHandler)

	server := &http.Server{
		Addr:              "localhost" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Second * 5,
	}

	log.Printf("Listening on http://localhost%s\n", port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
