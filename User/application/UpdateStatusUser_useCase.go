package application

import "github.com/bchanona/profile_warmheart_backend/User/domain"

type UpdateStatusUseCase struct {
    Repo domain.UserRepository
}

func NewUpdateStatusUseCase(repo domain.UserRepository) *UpdateStatusUseCase {
    return &UpdateStatusUseCase{Repo: repo}
}

func (uc *UpdateStatusUseCase) UpdateStatus(id int, premium bool) error {
    return uc.Repo.UpdateStatus(id, premium)
}