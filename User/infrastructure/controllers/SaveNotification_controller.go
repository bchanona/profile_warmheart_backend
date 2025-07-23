package controllers

import (

	"github.com/bchanona/profile_warmheart_backend/User/application"
	"github.com/bchanona/profile_warmheart_backend/User/domain"
	"github.com/gin-gonic/gin"
)

type SaveNotificationController struct {
	useCase *application.SaveNotificationUseCase
}

func NewSaveNotificationController(useCase *application.SaveNotificationUseCase)*SaveNotificationController{
	return &SaveNotificationController{useCase: useCase}
}

func (controller *SaveNotificationController) Execute(ctx *gin.Context){
	var notification domain.SaveNotifications

	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}
	err := controller.useCase.Execute(notification)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save notification"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Notification saved successfully"})
}