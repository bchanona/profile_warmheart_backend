package application

import "github.com/bchanona/profile_warmheart_backend/User/domain"

type UpdateUserUseCase struct {
	Repo domain.UserRepository
}

func NewUpdateUserUseCase (repo domain.UserRepository) *UpdateUserUseCase{
	return &UpdateUserUseCase{Repo: repo}
}

func (useCase *UpdateUserUseCase) Update(user domain.User) error{
	return useCase.Repo.Update(user)
}