package http

import (
	"net/http"

	"github.com/anucha-tk/go_todo/internal/application"
	"github.com/anucha-tk/go_todo/internal/templates/pages"
)

type TodoHandler struct {
	service *application.TodoService
}

func NewTodoHandler(service *application.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Home(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := pages.HomePage(todos).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
