package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dataninja-python/bookings/internal/config"
	"github.com/dataninja-python/bookings/internal/models"
	"github.com/dataninja-python/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
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
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals is the handler for the generals quarters room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// PostGenerals is the handler for the generals quarters room page
func (m *Repository) PostGenerals(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	_, err := w.Write([]byte(fmt.Sprintf("start date: %s and end date: %s", start, end)))
	if err != nil {
		log.Println("Error posting to search availability, ", err)
		return
	}
}

// Majors is the handler for the majors suite room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability is the handler for the majors suite room page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability is the handler for the majors suite room page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	_, err := w.Write([]byte(fmt.Sprintf("start date: %s and end date: %s", start, end)))
	if err != nil {
		log.Println("Error posting to search availability, ", err)
		return
	}
}

// jsonResponse type is used to test marshalling and unmarshalling required by go to manage json
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		return
	}
}

// Contact is the handler for the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is the handler for the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{})
}

// Reservation is the handler for the reservation page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	pkgAnnouncer()
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
