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
        UserID  int  `json:"user_id"`
        Premium bool `json:"premium"`
    }

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.UseCase.UpdateStatus(request.UserID, request.Premium); err != nil {
        if err == domain.ErrUserNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User status updated successfully"})
}