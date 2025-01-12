package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/smudgy-g/bookings/pkg/config"
	"github.com/smudgy-g/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the configuration for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds any data available to all templates
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using the template cache
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// get the template cache from the app config
	var templateCache map[string]*template.Template
	if app.ProductionMode {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatalf("Template %s not found in cache", tmpl)
	}

	// Create a buffer to store the rendered template
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	// Execute the template into the buffer
	_ = t.Execute(buf, td)

	// Write the buffer content to the response writer
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to response:", err)
	}
}

// createTemplateCache creates a map of pre-parsed templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// get all *.page.tmpl files
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	// iterate over pages and parse templates
	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the page file
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		// Get layouts and parse them into the current template set
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			// Add layouts to the existing template set
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		// Store the template set in the cache
		cache[name] = ts
	}

	return cache, nil
}
