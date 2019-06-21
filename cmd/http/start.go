// +build !local

package main

import (
	"net/http"
	"os"

	"bitbucket.org/ifood/goservice/cmd/http/services"

	log "github.com/inconshreveable/log15"
	newrelic "github.com/newrelic/go-agent"
	"github.com/spf13/viper"
)

func startApp(config *viper.Viper, logger log.Logger, userService *services.UserService, pushService *services.PushService) {
	host := config.GetString("host")
	appName := config.GetString("appname")
	nrLicense := config.GetString("new_relic_license")
	nrConfig := newrelic.NewConfig(appName, nrLicense)
	nrApp, err := newrelic.NewApplication(nrConfig)
	if err != nil {
		logger.Error("error creating New Relic Application", "err", err)
		os.Exit(1)
	}
	router := NewRouter(nrApp, logger, userService, pushService)
	server := &http.Server{
		Addr:    host,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		logger.Error("error starting server", "host", config.GetString("host"), "error", err)
		os.Exit(1)
	}
}
