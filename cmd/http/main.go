package main

import (
	"fmt"

	"bitbucket.org/ifood/goservice/cmd/http/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// initialize Viper configuration
	config := NewViperConfig()

	// initialize logger
	rootlog := RootLogger(config)

	// initialize database
	db, err := sqlx.Connect("postgres", config.GetString("datasource"))
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
