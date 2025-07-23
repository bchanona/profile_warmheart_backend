package controllers

import (
	"net/http"

	"github.com/bchanona/profile_warmheart_backend/User/application"
	"github.com/bchanona/profile_warmheart_backend/User/domain"
	"github.com/gin-gonic/gin"
)

type UpdateStatusController struct {
	UseCase *application.UpdateStatusUseCase
}

func NewUpdateStatusController(useCase *application.UpdateStatusUseCase) *UpdateStatusController {
	return &UpdateStatusController{UseCase: useCase}
}

func (ctrl *UpdateStatusController) UpdateStatus(c *gin.Context) {
	var request struct {
		Premium bool `json:"premium"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o faltante"})
		return
	}

	userID, ok := userIDRaw.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo interpretar el user_id"})
		return
	}

	if err := ctrl.UseCase.UpdateStatus(userID, request.Premium); err != nil {
		if err == domain.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User status updated successfully"})
}
