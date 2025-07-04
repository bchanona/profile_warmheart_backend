package domain

import "errors"

var (
    ErrUserAlreadyExists = errors.New("user already exists")
    ErrUserNotFound      = errors.New("user not found")
    ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidInput = errors.New("select all imput")
)

type UserRepository interface {
    Save(user User) error
    GetByEmail(email string) (User, error)
    GetByID(id int) (User, error)
    GetAll() ([]User, error)
    UpdateStatus(id int, premium bool) error
    Delete(id int) error
}