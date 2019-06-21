package main

import (
	"fmt"
	"net/http"

	"github.com/nubunto/go-boilerplate/cmd/http/services"
	"github.com/go-chi/render"
	log "github.com/inconshreveable/log15"
)

func listUsersEndpoint(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.FetchAll()
		if err != nil {
			render.Render(w, r, ErrRender(err, http.StatusInternalServerError))
			return
		}
		if err := render.RenderList(w, r, NewUserListResponse(users)); err != nil {
			render.Render(w, r, ErrRender(err, http.StatusUnprocessableEntity))
			return
		}
	}
}

func healthEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "OK")
		if err != nil {
			render.Render(w, r, ErrRender(err, http.StatusInternalServerError))
		}
	}
}

func snsPushEndpoint(logger log.Logger, pushService *services.PushService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use your own typed message here instead of a map
		// e.g. type MyMessage struct {...}
		// more efficient and type safe
		message := map[string]interface{}{
			"your": "message here",
		}
		mid, err := pushService.PushMessage(message)
		if err != nil {
			render.Render(w, r, ErrRender(err, http.StatusInternalServerError))
		}
		logger.Debug("sent message, got message id %v", mid)
	}
}
