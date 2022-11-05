package main

import (
	"fmt"
	"github.com/rostis232/booking/internal/config"
	"github.com/rostis232/booking/internal/handlers"
	"github.com/rostis232/booking/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

//const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application.
func main() {
	//Change app.InProduction to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", config.PortNumber)

	srv := &http.Server{
		Addr:    config.PortNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
