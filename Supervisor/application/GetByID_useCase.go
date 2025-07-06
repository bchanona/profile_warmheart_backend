package application

import domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"

type GetByIDUseCase struct {
	db domain.UserRepository
}

func GetByID(db domain.UserRepository) *GetByIDUseCase {
	return &GetByIDUseCase{db: db}
}
func (cp *GetByIDUseCase) Execute(id int32) (domain.Supervisor, error) {
	return cp.db.GetByIDSupervisor(id)
}
