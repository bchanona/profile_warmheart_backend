package domain

import "errors"

var (
	ErrSupervisorAlreadyExists = errors.New("supervisor already exists")
	ErrSupervisorNotFound      = errors.New("supervisor not found")
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrInvalidInput            = errors.New("select all imput")
	ErrNoSupervisorsFound      = errors.New("no supervisors found for the given user ID")
	ErrUserNotFound            = errors.New("user not found for given supervisor ID")
	ErrEmailAlreadyExists      = errors.New("email already in use")
)

type UserRepository interface { //luego añado los gets
	Save(supervisor Supervisor) error
	GetByIDSupervisor(id int32) (Supervisor, error)
	GetSupervisorsByUserID(id int32) ([]GetSupervisorByUserID, error)
	GetUserBySupervisorID(id int32) (GetUserBySupervisorID, error)
	UpdateSupervisor(id int, data UpdateSupervisor) error
	GetSupervisorByEmail(email string) (Supervisor, error)
	UpdateSupervisorPassword(id int, data UpdatePassword) error
	DeleteSupervisor(supervisorID int, userID int) error
	LoginSupervisors(emmail string, password string) (Supervisor, error)
}
