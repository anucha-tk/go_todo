package http

import (
	"strconv"

	"github.com/anucha-tk/go_todo/internal/application"
	"github.com/anucha-tk/go_todo/internal/common"
	"github.com/anucha-tk/go_todo/internal/domain"
	"github.com/anucha-tk/go_todo/internal/models"
	"github.com/anucha-tk/go_todo/internal/templates"
	"github.com/anucha-tk/go_todo/internal/templates/pages"
	"github.com/anucha-tk/go_todo/internal/templates/partials/tbody"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	service *application.TodoService
}

func NewTodoHandler(service *application.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Home(c *fiber.Ctx) error {
	todos, err := h.service.FindAll()
	if err != nil {
		return common.ResponseError(c, fiber.StatusInternalServerError, "Something went worng!", err)
	}

	return templates.Render(c, pages.HomePage(todos))
}

func (h *TodoHandler) Create(c *fiber.Ctx) error {
	var error error
	payload := new(models.TodoInput)
	error = c.BodyParser(&payload)
	if error != nil {
		return common.ResponseError(c, fiber.StatusInternalServerError, "Error Body Parser", error)
	}
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return common.ResponseError(c, fiber.StatusUnprocessableEntity, "Error invalid body", errors)
	}
	todo := domain.Todo{
		Description: payload.Description,
	}
	_, error = h.service.Create(todo)
	if error != nil {
		return common.ResponseError(c, fiber.StatusInternalServerError, "Something went worng!", errors)
	}

	todos, err := h.service.FindAll()
	if err != nil {
		return common.ResponseError(c, fiber.StatusInternalServerError, "Failed to fetch todos", err)
	}

	return templates.Render(c, tbody.TableBody(todos))
}

func (h *TodoHandler) Delete(c *fiber.Ctx) error {
	todoID := c.Params("id")
	id, err := strconv.ParseUint(todoID, 10, 32)
	if err != nil {
		return common.ResponseError(c, fiber.StatusInternalServerError, "ParseUint error", err)
	}
	err = h.service.Delete(uint(id))
	if err != nil {
		return common.ResponseError(c, fiber.StatusInternalServerError, "Delet error", err)
	}
	todos, err := h.service.FindAll()
	if err != nil {
		return common.ResponseError(c, fiber.StatusInternalServerError, "Failed to fetch todos", err)
	}

	return templates.Render(c, tbody.TableBody(todos))
}
