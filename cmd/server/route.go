package main

import (
	"github.com/anucha-tk/go_todo/internal/interfaces/http"
	"github.com/anucha-tk/go_todo/internal/templates"
	"github.com/anucha-tk/go_todo/internal/templates/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func routes(app *fiber.App, todoHandler *http.TodoHandler, aboutHandler *http.AboutHandler) {
	app.Use(logger.New())

	app.Get("/", todoHandler.Home)
	app.Post("/todos", todoHandler.Create)
	app.Delete("/todos/:id", todoHandler.Delete)
	app.Get("/about", aboutHandler.About)
	app.Use(notFoundMiddleware)
}

func notFoundMiddleware(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return templates.Render(c, pages.Notfound())
}
