package application

import domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"

type LoginUseCase struct {
	db domain.UserRepository
}

func NewLoginUseCase(db domain.UserRepository) *LoginUseCase {
	return &LoginUseCase{db: db}
}

func (useCase *LoginUseCase) Execute(email string, password string) (domain.Supervisor, error) {
	return useCase.db.LoginSupervisors(email, password)
}
