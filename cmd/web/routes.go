package main

import (
	"github.com/dataninja-python/bookings/pkg/config"
	"github.com/dataninja-python/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// routes sets the routes for the application
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(middleware.Logger)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	// return handlers
	return mux
}
