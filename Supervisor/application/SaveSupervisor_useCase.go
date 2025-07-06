package application

import domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"

type SaveSupervisorUseCase struct {
	db domain.UserRepository
}

func SaveSupervisor(db domain.UserRepository) *SaveSupervisorUseCase {
	return &SaveSupervisorUseCase{db: db}
}
func (cp *SaveSupervisorUseCase) Execute(Supervisor domain.Supervisor) error {
	if Supervisor.Name == "" || Supervisor.Surnames == "" || Supervisor.Email == "" || Supervisor.Password == "" || Supervisor.User_id == 0 {
		return domain.ErrInvalidInput
	}
	return cp.db.Save(Supervisor)
}
