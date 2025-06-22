package domain

import "errors"

var (
	ErrSupervisorAlreadyExists = errors.New("supervisor already exists")
	ErrSupervisorNotFound      = errors.New("supervisor not found")
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrInvalidInput            = errors.New("select all imput")
)

type UserRepository interface { //luego añado los gets
	Save(supervisor Supervisor) error
}
