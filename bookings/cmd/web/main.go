package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/driver"

	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/helpers"

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
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = db.SQL.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app), //nolint:typecheck
	}

	log.Println("server is running", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=isoppp password=")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
		return nil, err
	}
	log.Println("Connected to database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return db, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	_ = fmt.Sprintf("Staring application on port %s\n", portNumber)

	helpers.NewHelpers(&app)
	return db, nil
}
