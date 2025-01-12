package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuration
type AppConfig struct {
	ProductionMode bool
	TemplateCache  map[string]*template.Template
	SessionManager *scs.SessionManager
}
