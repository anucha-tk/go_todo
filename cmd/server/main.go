package main

import (
	"flag"
	"log"

	"github.com/anucha-tk/go_todo/internal/application"
	"github.com/anucha-tk/go_todo/internal/infrastructure/db"
	httphandler "github.com/anucha-tk/go_todo/internal/interfaces/http"
	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()
	app.Static("/", "./assets/")

	routes(app, todoHandler, aboutHandler)

	log.Fatal(app.Listen(":3000"))
}
