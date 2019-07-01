// +build local

package main

import (
	"net/http"
	"os"

	log "github.com/inconshreveable/log15"
	"github.com/spf13/viper"
)

func startApp(config *viper.Viper, logger log.Logger, env *Env) {
	host := config.GetString("host")
	logger.Info("starting local HTTP server", "host", host)
	router := NewRouter(nil, logger, env)
	server := &http.Server{
		Addr:    host,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Error("error starting server", "host", config.GetString("host"), "error", err)
		os.Exit(1)
	}
}
