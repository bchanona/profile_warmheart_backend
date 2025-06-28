package controllers

import (
	"errors"
	"net/http"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/gin-gonic/gin"
)

type GetUserBySupervisorIDController struct {
	useCase *application.GetUserBySupervisorIDUseCase
}

func NewGetUserBySupervisorIDController(useCase *application.GetUserBySupervisorIDUseCase) *GetUserBySupervisorIDController {
	return &GetUserBySupervisorIDController{useCase: useCase}
}

func (c *GetUserBySupervisorIDController) Execute(ctx *gin.Context) {
	idRaw, exists := ctx.Get("supervisor_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}
	supervisorID := idRaw.(int)

	user, err := c.useCase.Execute(int32(supervisorID))
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found for supervisor"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
