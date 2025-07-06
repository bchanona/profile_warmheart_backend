package application

import domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"

type GetUserBySupervisorIDUseCase struct {
	repo domain.UserRepository
}

func NewGetUserBySupervisorIDUseCase(repo domain.UserRepository) *GetUserBySupervisorIDUseCase {
	return &GetUserBySupervisorIDUseCase{repo: repo}
}

func (uc *GetUserBySupervisorIDUseCase) Execute(id int32) (domain.GetUserBySupervisorID, error) {
	user, err := uc.repo.GetUserBySupervisorID(id)
	if err != nil {
		return domain.GetUserBySupervisorID{}, err
	}
	return user, nil
}
