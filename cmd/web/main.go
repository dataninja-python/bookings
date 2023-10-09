package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/dataninja-python/bookings/pkg/config"
	"github.com/dataninja-python/bookings/pkg/handlers"
	"github.com/dataninja-python/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

// variables created here are available throughout the packages
// variable storing the desired server port
const portNumber = ":8080"

// variable stores application configuration and makes it available throughout this package
var app config.AppConfig

// variable stores a pointer to the cookie manager pointer
var session *scs.SessionManager

// main is the main function
func main() {
	// setting app.InProduction variable changes several parts of the application to production
	//
	// changes include (but should be routinely checked and updated):
	// - making session cookies secure
	app.InProduction = false
	log.Println("Application running in production? ", app.InProduction)

	// manage sessions
	session = scs.New()
	session.Lifetime = 12 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// store the session into the config structure
	app.Session = session

	// create variable to store template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// store the template in the app variable
	app.TemplateCache = tc
	// set whether to use the cache
	//
	// false: render templates when routes are request
	// true: use the cache when routes are requested
	app.UseCache = false

	// create a variable that stores a repository
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// render new templates
	render.NewTemplates(&app)

	// reporting that the server is starting
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	// launches server with provided port and handler
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// execute the servers
	err = server.ListenAndServe()
	log.Fatal(err)
}
