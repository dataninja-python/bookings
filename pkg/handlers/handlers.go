package handlers

import (
	"github.com/dataninja-python/bookings/pkg/config"
	"github.com/dataninja-python/bookings/pkg/models"
	"github.com/dataninja-python/bookings/pkg/render"
	"log"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// pkgAnnouncer logs that work performed here is from this package
func pkgAnnouncer() {
	pkg := "handler"
	log.Printf("The program is in the %s package\n", pkg)
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	pkgAnnouncer()
	log.Println("Creating a new Repository object")
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	pkgAnnouncer()
	log.Println("Assigns value to package wide Repository variable")
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	// variable stores the IP address of visitor
	remoteIP := r.RemoteAddr
	// every time a user visits this route the ip address is stored
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	log.Println("Returning rendered home.page.tmpl")
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()

	// some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// gets the IP address form user visiting home
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	log.Println("Returning rendered home.page.tmpl")
	// render the templates to the browser
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals is the handler for the generals quarters room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// Majors is the handler for the majors suite room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// Availability is the handler for the majors suite room page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact is the handler for the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is the handler for the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, "reservation.page.tmpl", &models.TemplateData{})
}

// Reservation is the handler for the reservation page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}
