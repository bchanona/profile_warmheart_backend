package application

import "github.com/bchanona/profile_warmheart_backend/User/domain"

type GetByIDUserUseCase struct {
	Repo domain.UserRepository
}

func NewGetByIDUserUseCase(repo domain.UserRepository) *GetByIDUserUseCase {
	return &GetByIDUserUseCase{Repo: repo}
}

func (useCase *GetByIDUserUseCase) GetByID(user_id int) (domain.User, error){
	return useCase.Repo.GetByID(user_id)
} 