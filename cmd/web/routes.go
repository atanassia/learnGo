package main

import (
	"net/http"

	"github.com/atanassia/go-course/pkg/config"
	"github.com/atanassia/go-course/pkg/handlers"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurve)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
