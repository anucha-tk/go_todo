package main

import (
	"github.com/anucha-tk/go_todo/internal/assets"
	"github.com/anucha-tk/go_todo/internal/interfaces/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(todoHandler *http.TodoHandler, aboutHandler *http.AboutHandler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", todoHandler.Home)
	router.Get("/about", aboutHandler.About)
	assets.Mount(router)
	return router
}
