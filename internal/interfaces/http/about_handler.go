package http

import (
	"github.com/anucha-tk/go_todo/internal/templates"
	"github.com/anucha-tk/go_todo/internal/templates/pages"
	"github.com/gofiber/fiber/v2"
)

type AboutHandler struct{}

func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

func (h *AboutHandler) About(c *fiber.Ctx) error {
	return templates.Render(c, pages.AboutPage())
}
