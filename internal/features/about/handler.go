package about

import (
	"net/http"

	"github.com/anucha-tk/go_todo/internal/templates/pages"
	"github.com/go-chi/chi/v5"
)

type (
	Handler interface {
		About(w http.ResponseWriter, r *http.Request)
	}
	handler struct{}
)

func NewHandler() Handler {
	return &handler{}
}

func Mount(r chi.Router, h Handler) {
	r.Get("/about", h.About)
}

func (h *handler) About(w http.ResponseWriter, r *http.Request) {
	if err := pages.AboutPage().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
