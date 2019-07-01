// This file contains interfaces to services defined in cmd/http/services.
// Why? 'cause it's easy to test with interfaces than it is to test with concrete implementations.
// Also, when you add new functionality in your service, it becomes clear where you need to implement
// in your handlers.
package main

import goservice "github.com/nubunto/go-boilerplate"

// UserService handles user related things
type UserService interface {
	FetchAll() ([]goservice.User, error)
}

// PushService handles push related things
type PushService interface {
	PushMessage(interface{}) (string, error)
}

// Env holds all your services in a struct.
// This exists as a helper to facilitate distributing your services to handlers
type Env struct {
	UserService UserService
	PushService PushService
}
