package main

import (
	log "github.com/inconshreveable/log15"
	"github.com/spf13/viper"
)

// RootLogger returns a root log15 logger
// that can be customized and derived from
func RootLogger(config *viper.Viper) log.Logger {
	appName := config.Get("appname")
	env := config.Get("environment")
	return log.New("app", appName, "env", env)
}
