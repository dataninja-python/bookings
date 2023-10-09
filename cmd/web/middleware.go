package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// WriteToConsole writes to the console as an example of writing custom middleware
func WriteToConsole(next http.Handler) http.Handler {
	// returns a handler as an anonymous function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")

		next.ServeHTTP(w, r)
	})
}

// NoSurf protects users from cross-site requirest forgery (CSRF)
//
// without this a bad actor can trick the application into taking
// malicious actions using session storage data of an authenticated user
// NOTE: this is a critical basic security feature in applications
// that use client-server architecture where communication comes from the front-end
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
