package controllers

import (
	"errors"
	"net/http"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordController struct {
	useCase *application.UpdateSupervisorPasswordUseCase
}

func NewUpdatePasswordController(useCase *application.UpdateSupervisorPasswordUseCase) *UpdatePasswordController {
	return &UpdatePasswordController{useCase: useCase}
}

func (c *UpdatePasswordController) Execute(ctx *gin.Context) {
	idRaw, exists := ctx.Get("supervisor_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}
	id := idRaw.(int)

	var data domain.UpdatePassword
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := c.useCase.Execute(id, data)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidInput):
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		case errors.Is(err, domain.ErrSupervisorNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": "supervisor not found"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update password"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}
