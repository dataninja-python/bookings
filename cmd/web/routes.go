package main

import (
	"github.com/dataninja-python/bookings/pkg/config"
	"github.com/dataninja-python/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// routes sets the routes for the application
func routes(app *config.AppConfig) http.Handler {
	// create a new router
	mux := chi.NewRouter()

	// --- middleware --- //
	// server recovers gracefully from a crash
	mux.Use(middleware.Recoverer)
	// mux.Use(middleware.Logger)
	// a personally written middleware
	mux.Use(WriteToConsole)
	// CSRF protection
	mux.Use(NoSurf)
	// manages session interactions
	mux.Use(SessionLoad)

	// listen for the routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// get and store the image file from the static directory
	fileServer := http.FileServer(http.Dir("./static/"))
	// render the image to the browser
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// return handlers
	return mux
}
