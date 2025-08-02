package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hernanvza/bookings/internal/config"
	"github.com/Hernanvza/bookings/internal/models"
	"github.com/Hernanvza/bookings/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform logic business
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// General renders the make a reservation page and displays the form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// General renders the make a reservation page and displays the form
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Major renders the make a reservation page and displays the form
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the make a reservation page and displays the form
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page and processes the form submission
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type JSONResponse struct {
	Start   string `json:"start"`
	End     string `json:"end"`
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON renders the availability as JSON
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	resp := JSONResponse{
		Start:   r.Form.Get("start"),
		End:     r.Form.Get("end"),
		OK:      true,
		Message: "Availability received",
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		json.NewEncoder(w).Encode(resp)
		return
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the contact page and displays the form
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
