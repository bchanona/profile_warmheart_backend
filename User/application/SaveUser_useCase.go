package application

import "github.com/bchanona/profile_warmheart_backend/User/domain"

type SaveUserUseCase struct {
	Repo domain.UserRepository
}

func NewSaveUserUseCase(repo domain.UserRepository) *SaveUserUseCase{
	return &SaveUserUseCase{Repo: repo}
}

func (useCase *SaveUserUseCase) Save(user domain.User) error{
	if user.Name == "" || user.Email == "" || user.Password == "" {
        return domain.ErrInvalidInput
    }
	return useCase.Repo.Save(user)
}