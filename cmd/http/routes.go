package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"

	"github.com/nubunto/go-boilerplate/cmd/http/services"

	log "github.com/inconshreveable/log15"

	newrelic "github.com/newrelic/go-agent"
)

// NewRouter returns a router that contains this project's
// HTTP routes.
func NewRouter(app newrelic.Application, logger log.Logger, userService *services.UserService, pushService *services.PushService) *chi.Mux {
	router := chi.NewRouter()
	// middleware for all routes
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// This is where you add your routes
	router.Route("/users", func(r chi.Router) {
		// middleware for /users route only
		// r.Use(...)
		// GET /users/, instrumented by New Relic
		r.Get(WithNewRelic(app, "/", listUsersEndpoint(userService)))
		r.Post(WithNewRelic(app, "/", snsPushEndpoint(logger, pushService)))
	})

	// GET /health, not instrumented by New Relic
	router.Get("/health", healthEndpoint())

	return router
}

func WithNewRelic(app newrelic.Application, pattern string, handler http.HandlerFunc) (string, http.HandlerFunc) {
	if app == nil {
		// do not wrap if no Application is available
		// which is the case when running locally
		return pattern, handler
	}
	return newrelic.WrapHandleFunc(app, pattern, handler)
}
