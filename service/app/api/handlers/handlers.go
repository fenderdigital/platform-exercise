package handlers

import (
	"log"
	"net/http"
	"os"

	"platform-exercise/service/business/auth"
	"platform-exercise/service/business/mid"
	"platform-exercise/service/foundation/web"

	"github.com/jmoiron/sqlx"
)

// API constructs an http.Handler with all application routes defined.
func API(shutdown chan os.Signal, log *log.Logger, db *sqlx.DB, a *auth.Auth) http.Handler {

	// Construct the web.App which holds all routes as well as common Middleware.
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics(log))

	// Register health check endpoint. This route is not authenticated.
	c := check{
		db: db,
	}
	app.Handle(http.MethodGet, "/", c.landing)
	app.Handle(http.MethodGet, "/health", c.health)

	o := operHandlers{
		db:   db,
		auth: a,
	}

	app.Handle(http.MethodPost, "/register", o.register, mid.Authenticate(a), mid.Authorized(auth.RoleAdmin))
	app.Handle(http.MethodPost, "/login", o.login)

	return app
}
