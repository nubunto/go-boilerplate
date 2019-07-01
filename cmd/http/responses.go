package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	goservice "github.com/nubunto/go-boilerplate"
)

func RenderJSON(w http.ResponseWriter, code int, data interface{}) error {
	m, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}
	w.WriteHeader(code)
	_, err = w.Write(m)
	if err != nil {
		return fmt.Errorf("error writing response to wire: %v", err)
	}
	return nil
}

type UserResponse struct {
	User goservice.User `json:"user"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

func NewUsersResponse(users []goservice.User) *UsersResponse {
	r := &UsersResponse{
		Users: make([]UserResponse, len(users)),
	}
	for i, u := range users {
		r.Users[i] = UserResponse{u}
	}
	return r
}

func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{
		Err:       err,
		ErrorText: err.Error(),
	}
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	ErrorText string `json:"error"`
}
