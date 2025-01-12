package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/smudgy-g/bookings/pkg/config"
	"github.com/smudgy-g/bookings/pkg/handlers"
	"github.com/smudgy-g/bookings/pkg/render"
)

const portNumber = ":42069"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.SessionManager = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache:", err)
	}

	app.TemplateCache = tc
	app.ProductionMode = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("server listening on port %s\n", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
