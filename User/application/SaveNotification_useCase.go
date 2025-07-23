package application

import "github.com/bchanona/profile_warmheart_backend/User/domain"

type SaveNotificationUseCase struct {
	db domain.UserRepository
}

func NewSaveNotificationUseCase(db domain.UserRepository)*SaveNotificationUseCase{
	return &SaveNotificationUseCase{db: db}
}

func (uc *SaveNotificationUseCase) Execute(notification domain.SaveNotifications) error{
	return uc.db.SaveNotification(notification)
}