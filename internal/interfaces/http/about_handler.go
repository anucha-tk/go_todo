package http

import (
	"net/http"

	"github.com/anucha-tk/go_todo/internal/templates/pages"
)

type AboutHandler struct{}

func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

func (h *AboutHandler) About(w http.ResponseWriter, r *http.Request) {
	if err := pages.AboutPage().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
