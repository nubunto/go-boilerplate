package main

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

func NewViperConfig() *viper.Viper {
	config := viper.New()

	setConfigDefaults(config)
	// could use a config file, if you want
	// setConfigFile(config)
	setEnvironmentConfig(config)

	return config
}

func setConfigDefaults(config *viper.Viper) {
	config.SetDefault("host", "0.0.0.0:8080")
	config.SetDefault("appname", "goservice")
	config.SetDefault("environment", "local")
	config.SetDefault("datasource", "postgres://postgres:postgres@localhost:5432/goservice?sslmode=disable")
}

// example code for Viper configuration files
func setConfigFile(config *viper.Viper) {
	appName := config.GetString("appname")
	config.SetConfigName("application")
	config.AddConfigPath(path.Join("/etc", appName))
	config.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		// panic on bad config
		panic(fmt.Errorf("error reading configuration: %v", err))
	}
}

func setEnvironmentConfig(config *viper.Viper) {
	appName := config.GetString("appname")
	config.SetEnvPrefix(appName)
	config.AutomaticEnv()
}
