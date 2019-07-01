package main

import (
	goservice "github.com/nubunto/go-boilerplate"
	"github.com/nubunto/go-boilerplate/cmd/http/services"
)

type inMemoryUserService struct {
	users []goservice.User
}

func (s *inMemoryUserService) FetchAll() ([]goservice.User, error) {
	return s.users, nil
}

func newInMemoryUserService() *inMemoryUserService {
	return &inMemoryUserService{
		users: make([]goservice.User, 0, 100),
	}
}

func main() {
	// initialize Viper configuration
	config := NewViperConfig()

	// initialize logger
	rootlog := RootLogger(config)

	// initialize real database
	// db, err := sqlx.Open("postgres", config.GetString("datasource"))
	// if err != nil {
	// 	panic(fmt.Errorf("error connecting to postgres: %v", err))
	// }

	// initialize in-memory user service

	// initialize services as needed
	awsSession := NewAWSSession()
	snsPublisher := NewSNS(awsSession)
	snsExampleService := services.NewPushService(config.GetString("your_topic_arn"), snsPublisher)
	//userService := services.NewUserService(db)
	userService := newInMemoryUserService()

	env := &Env{
		UserService: userService,
		PushService: snsExampleService,
	}

	// start your application, depending on your environment
	startApp(config, rootlog, env)
}
