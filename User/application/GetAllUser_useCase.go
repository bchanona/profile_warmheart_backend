package application

import "github.com/bchanona/profile_warmheart_backend/User/domain"

type GetAllUsersUseCase struct {
    Repo domain.UserRepository
}

func NewGetAllUsersUseCase(repo domain.UserRepository) *GetAllUsersUseCase {
    return &GetAllUsersUseCase{Repo: repo}
}

func (uc *GetAllUsersUseCase) GetAll() ([]domain.User, error) {
    return uc.Repo.GetAll()
}