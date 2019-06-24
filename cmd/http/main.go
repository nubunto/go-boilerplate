package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nubunto/go-boilerplate/cmd/http/services"
)

func main() {
	// initialize Viper configuration
	config := NewViperConfig()

	// initialize logger
	rootlog := RootLogger(config)

	// initialize database
	db, err := sqlx.Open("postgres", config.GetString("datasource"))
	if err != nil {
		panic(fmt.Errorf("error connecting to postgres: %v", err))
	}

	// initialize services as needed
	awsSession := NewAWSSession()
	snsPublisher := NewSNS(awsSession)
	snsExampleService := services.NewPushService(config.GetString("your_topic_arn"), snsPublisher)
	userService := services.NewUserService(db)

	// start your application, depending on your environment
	startApp(config, rootlog, userService, snsExampleService)
}
