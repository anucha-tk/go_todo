package main

import (
	"github.com/anucha-tk/go_todo/internal/assets"
	"github.com/anucha-tk/go_todo/internal/features/home"
	"github.com/go-chi/chi/v5"
)

func routes() *chi.Mux {
	router := chi.NewRouter()
	home.Mount(router, home.NewHandler())
	assets.Mount(router)
	return router
}
