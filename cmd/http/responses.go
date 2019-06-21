package main

import (
	"net/http"

	"bitbucket.org/ifood/goservice"
	"github.com/go-chi/render"
)

type UserResponse struct {
	User goservice.User
}

func (ur *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserListResponse(users []goservice.User) []render.Renderer {
	list := make([]render.Renderer, len(users))
	for _, user := range users {
		list = append(list, &UserResponse{User: user})
	}
	return list
}

func ErrRender(err error, code int) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: code,
		ErrorText:      err.Error(),
	}
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	ErrorText  string `json:"error"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	e.StatusText = http.StatusText(e.HTTPStatusCode)
	return nil
}
