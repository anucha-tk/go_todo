package home

import (
	"net/http"

	"github.com/anucha-tk/go_todo/internal/templates/pages"
	"github.com/go-chi/chi/v5"
)

type (
	Handler interface {
		Home(w http.ResponseWriter, r *http.Request)
	}
	handler struct{}
)

func NewHandler() Handler {
	return &handler{}
}

func Mount(r chi.Router, h Handler) {
	r.Get("/", h.Home)
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	if err := pages.HomePage().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
