package controllers

import (
	"errors"
	"net/http"
	"strconv"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/gin-gonic/gin"
)

type UpdateSupervisorController struct {
	useCase *application.UpdateSupervisorUseCase
}

func NewUpdateSupervisorController(useCase *application.UpdateSupervisorUseCase) *UpdateSupervisorController {
	return &UpdateSupervisorController{useCase: useCase}
}

func (c *UpdateSupervisorController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid supervisor ID"})
		return
	}

	var updateData domain.UpdateSupervisor
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = c.useCase.Execute(id, updateData)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidInput):
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
		case errors.Is(err, domain.ErrEmailAlreadyExists):
			ctx.JSON(http.StatusConflict, gin.H{"error": "email already in use"})
		case errors.Is(err, domain.ErrSupervisorNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": "supervisor not found"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update supervisor"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "supervisor updated successfully"})
}
