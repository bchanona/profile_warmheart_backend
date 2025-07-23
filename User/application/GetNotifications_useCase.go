package application

import "github.com/bchanona/profile_warmheart_backend/User/domain"

type GetNotificationUseCase struct {
	db domain.UserRepository
}

func NewGetNotificationUseCase(db domain.UserRepository) *GetNotificationUseCase{
	return &GetNotificationUseCase{db: db}
}

func (uc *GetNotificationUseCase) Execute(user_id int)([]domain.GetNotifications,error){
	return uc.db.GetNotification(user_id)
}
