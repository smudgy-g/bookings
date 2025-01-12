package handlers

import (
	"net/http"

	"github.com/smudgy-g/bookings/pkg/config"
	"github.com/smudgy-g/bookings/pkg/models"
	"github.com/smudgy-g/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// The rpository used by the handlers
var Repo *Repository

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.SessionManager.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["email"] = "adamgriff86@gmail.com"

	remoteIP := m.App.SessionManager.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
