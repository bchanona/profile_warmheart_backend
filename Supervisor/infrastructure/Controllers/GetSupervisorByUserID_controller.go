package controllers

import (
	"errors"
	"net/http"
	"strconv"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/gin-gonic/gin"
)

type GetSupervisorByUserIdController struct {
	useCase *application.GetSupervisorByUserIDUseCase
}

func NewGetSupervisorByUserIdController(useCase *application.GetSupervisorByUserIDUseCase) *GetSupervisorByUserIdController {
	return &GetSupervisorByUserIdController{useCase: useCase}
}

func (c *GetSupervisorByUserIdController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	supervisors, err := c.useCase.Execute(int32(id))
	if err != nil {
		if errors.Is(err, domain.ErrNoSupervisorsFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "no supervisors found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"supervisors": supervisors})
}
