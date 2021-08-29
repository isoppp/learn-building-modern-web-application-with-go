package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/models"

	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/config"
	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/handlers"
	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":5555"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app), //nolint:typecheck
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	_ = fmt.Sprintf("Staring application on port %s\n", portNumber)

	return nil
}
