package services

import (
	"github.com/nubunto/go-boilerplate"
	"github.com/jmoiron/sqlx"
)

type UserService struct {
	db *sqlx.DB
}

// NewUserService returns a new UserService compliant structure
// that reads user from a SQL database using the `sqlx` package.
func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (us *UserService) FetchAll() (users []goservice.User, err error) {
	if err := us.db.Select(&users, "SELECT name FROM users"); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) FetchByID(string) (goservice.User, error) {
	panic("not implemented")
}
func (us *UserService) UpdateByID(string, goservice.User) error {
	panic("not implemented")
}
func (us *UserService) DeleteByID(string) error {
	panic("not implemented")
}
