package application

import "github.com/bchanona/profile_warmheart_backend/User/domain"

type DeleteUserUseCase struct {
    Repo domain.UserRepository
}

func NewDeleteUserUseCase(repo domain.UserRepository) *DeleteUserUseCase {
    return &DeleteUserUseCase{Repo: repo}
}

func (uc *DeleteUserUseCase) Delete(id int) error {
    return uc.Repo.Delete(id)
}