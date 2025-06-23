package controllers

import (
	"errors"
	"net/http"
	"strconv"

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
	idParam := ctx.Param("id")
	supervisorID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid supervisor ID"})
		return
	}

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
