package application

import domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"

type UpdateSupervisorUseCase struct {
	db domain.UserRepository
}

func NewUpdateSupervisorUseCase(db domain.UserRepository) *UpdateSupervisorUseCase {
	return &UpdateSupervisorUseCase{db: db}
}

func (uc *UpdateSupervisorUseCase) Execute(id int, data domain.UpdateSupervisor) error {
	if data.Name == "" || data.Surnames == "" || data.Email == "" {
		return domain.ErrInvalidInput
	}

	// Verificar colisión de email
	existing, err := uc.db.GetSupervisorByEmail(data.Email)
	if err == nil && existing.Supervisor_id != id {
		return domain.ErrEmailAlreadyExists
	}

	// Intentar actualizar
	err = uc.db.UpdateSupervisor(id, data)
	if err != nil {
		return err
	}

	return nil
}
