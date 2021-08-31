package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/models"

	"github.com/alexedwards/scs/v2"
	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/config"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session
	app = &testApp

	os.Exit(m.Run())
}
