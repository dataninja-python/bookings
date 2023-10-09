package render

import (
	"bytes"
	"fmt"
	"github.com/dataninja-python/bookings/pkg/config"
	"github.com/dataninja-python/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// pkgAnnouncer logs that work performed here is from this package
func pkgAnnouncer() {
	pkg := "render"
	log.Printf("The program is in the %s package\n", pkg)
}

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	// log this is in the render package
	pkgAnnouncer()
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// log this is in the render package
	pkgAnnouncer()

	return td
}

// RenderTemplate renders a page from a template
//
// The function combines a base layout template with page specific details
// The final result is a new specific page presented as a view by the server
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// create err variable to hold errors throughout this function
	var err error

	// log this is in the render package
	pkgAnnouncer()

	// map variable that holds templates
	//
	// @key type: string
	// @value type: pointer to template.Template
	log.Println("map with string key and *template.Template value created")
	var tc map[string]*template.Template

	// make a conditional decision to use the cache
	//
	// @app.UseCache: boolean
	// @true: get and use the application cache
	// @false: render the template from disk each time
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		// render template from disk each time
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Println("error attempting to use template cache")
			log.Println("error occurs @package: render, @function: RenderTemplate, @section: app.UseCache " +
				"conditional")
			log.Println("ignoring desire to use cache and getting template cache from disk instead")
			tc = app.TemplateCache
		}
	}

	// search the template map for the desired template
	//
	// @result: template, bool
	// @exists: bool, found = true, not found = false
	// @t: template pointer if found, 0 if not found
	t, exists := tc[tmpl]

	// check if the key exists in the map
	//
	// @false: fatal error; log the message and end the program
	// @TODO: implement more graceful way of dealing with fatal errors
	// @TODO: preferably use pattern inspired by Elixir/Erlang where this function can fail
	// @TODO: the rest of the program continues running while this portion is archived for learning and rebooted
	if !exists {
		log.Fatal("Could not get template from template cache")
	}

	// create an empty buffer variable to store templates as bytes
	//
	// @why: ensures templates are only made available once fully rendered
	// @why: starts empty but can variably change size as needed, unless otherwise limited
	// @why: provides read, write, and other functionality
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	// transfer the template to the buffer
	err = t.Execute(buf, td)
	if err != nil {
		log.Println("error putting the template in the buffer")
	}

	// write template to browser from buffer
	var bytesWritten int64
	bytesWritten, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
	log.Println("Bytes written: ", bytesWritten)
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// create variable to store errors
	var err error

	// log this is in the render package
	pkgAnnouncer()

	// create map to store templates
	myCache := map[string]*template.Template{}

	// create a variable to store a string array
	var pages []string
	pages, err = filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for index, page := range pages {
		log.Println("The current index: ", index)
		// file name
		name := filepath.Base(page)
		// get the templates
		ts, _ := template.New(name).Funcs(functions).ParseFiles(page)
		// error
		_, err = template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("error creating the templates from files")
			log.Println(err)
			return myCache, err
		}

		// create the string array
		var matches []string
		// return the string array with all templates
		matches, err = filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// if there are templates deal with it
		if len(matches) > 0 {
			// combine page templates with layout templates
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// add the template to the cache
		log.Println("adding ", ts, "to template cache")
		myCache[name] = ts
	}
	// return the cache and error message
	return myCache, nil
}
