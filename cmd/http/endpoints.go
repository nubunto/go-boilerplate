package main

import (
	"fmt"
	"net/http"

	log "github.com/inconshreveable/log15"
)

func listUsersEndpoint(userService UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.FetchAll()
		if err != nil {
			RenderJSON(w, http.StatusInternalServerError, NewErrResponse(err))
			return
		}
		RenderJSON(w, http.StatusOK, NewUsersResponse(users))
	}
}

func healthEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "OK")
		if err != nil {
			RenderJSON(w, http.StatusInternalServerError, NewErrResponse(err))
		}
	}
}

func snsPushEndpoint(logger log.Logger, pushService PushService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use your own typed message here instead of a map
		// e.g. type MyMessage struct {...}
		// more efficient and type safe
		message := map[string]interface{}{
			"your": "message here",
		}
		mid, err := pushService.PushMessage(message)
		if err != nil {
			RenderJSON(w, http.StatusInternalServerError, NewErrResponse(err))
		}
		logger.Debug("sent message", "message id", mid)
	}
}
