package main

import (
	"github.com/anucha-tk/go_todo/internal/assets"
	"github.com/anucha-tk/go_todo/internal/features/about"
	"github.com/anucha-tk/go_todo/internal/features/home"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	home.Mount(router, home.NewHandler())
	about.Mount(router, about.NewHandler())
	assets.Mount(router)
	return router
}
