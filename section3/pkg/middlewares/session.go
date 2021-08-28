package middlewares

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/config"
)

var session *scs.SessionManager

func init() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = config.GetConfig().InProduction
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
