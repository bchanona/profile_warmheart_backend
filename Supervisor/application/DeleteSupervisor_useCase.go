package application

import domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"

type DeleteSupervisorUseCase struct {
	repo domain.UserRepository
}

func NewDeleteSupervisorUseCase(repo domain.UserRepository) *DeleteSupervisorUseCase {
	return &DeleteSupervisorUseCase{repo: repo}
}

func (uc *DeleteSupervisorUseCase) Execute(id int) error {
	return uc.repo.DeleteSupervisor(id)
}
