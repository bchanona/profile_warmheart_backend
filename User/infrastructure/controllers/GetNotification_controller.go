package controllers

import (
	"net/http"

	"github.com/bchanona/profile_warmheart_backend/User/application"
	"github.com/gin-gonic/gin"
)

type GetNotificationController struct {
	useCase *application.GetNotificationUseCase
}

func NewGetNotificationController(useCase *application.GetNotificationUseCase) *GetNotificationController {
	return &GetNotificationController{useCase: useCase}
}

func (controller *GetNotificationController) Execute(ctx *gin.Context) {

	userIdstr, exists := ctx.Get("user_id")

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id de usuario no proporcionado"})
	}

	userId, ok := userIdstr.(int)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}
	notifications, err := controller.useCase.Execute(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Notifications": notifications})

}
