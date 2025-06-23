package controllers

import (
	"errors"
	"net/http"
	"strconv"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/gin-gonic/gin"
)

type DeleteSupervisorController struct {
	useCase *application.DeleteSupervisorUseCase
}

func NewDeleteSupervisorController(useCase *application.DeleteSupervisorUseCase) *DeleteSupervisorController {
	return &DeleteSupervisorController{useCase: useCase}
}

func (c *DeleteSupervisorController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid supervisor ID"})
		return
	}

	err = c.useCase.Execute(id)
	if err != nil {
		if errors.Is(err, domain.ErrSupervisorNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "supervisor not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete supervisor"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "supervisor deleted successfully"})
}
