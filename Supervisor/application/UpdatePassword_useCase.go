package application

import domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"

type UpdateSupervisorPasswordUseCase struct {
	repo domain.UserRepository
}

func NewUpdateSupervisorPasswordUseCase(repo domain.UserRepository) *UpdateSupervisorPasswordUseCase {
	return &UpdateSupervisorPasswordUseCase{repo: repo}
}

func (uc *UpdateSupervisorPasswordUseCase) Execute(id int, data domain.UpdatePassword) error {
	if data.Password == "" {
		return domain.ErrInvalidInput
	}
	return uc.repo.UpdateSupervisorPassword(id, data)
}
