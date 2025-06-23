package application

import domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"

type GetSupervisorByUserIDUseCase struct {
	db domain.UserRepository
}

func NewGetSupervisorByUserIDUseCase(db domain.UserRepository) *GetSupervisorByUserIDUseCase {
	return &GetSupervisorByUserIDUseCase{db: db}
}
func (uc *GetSupervisorByUserIDUseCase) Execute(userID int32) ([]domain.GetSupervisorByUserID, error) {
	supervisors, err := uc.db.GetSupervisorsByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(supervisors) == 0 {
		return nil, domain.ErrNoSupervisorsFound
	}
	return supervisors, nil
}
