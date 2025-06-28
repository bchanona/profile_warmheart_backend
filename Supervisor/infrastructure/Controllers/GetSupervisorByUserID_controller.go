package controllers

import (
	"errors"
	"net/http"

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
	userIDRaw, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}
	userID := userIDRaw.(int)

	supervisors, err := c.useCase.Execute(int32(userID))
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
