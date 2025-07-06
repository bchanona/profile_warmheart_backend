package controllers

import (
	"net/http"
	"strconv"
	"time"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/gin-gonic/gin"
)

type GetByIdController struct {
	useCaseGetById *application.GetByIDUseCase
}

func NewGetByIdProductController(useCaseGetById *application.GetByIDUseCase) *GetByIdController {
	return &GetByIdController{useCaseGetById: useCaseGetById}
}

func (c *GetByIdController) Execute(ctx *gin.Context) {
	supervisorIDParam := ctx.Param("id")
	supervisorID, err := strconv.Atoi(supervisorIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid supervisor ID"})
		return
	}

	userIDRaw, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}
	userID := userIDRaw.(int)

	initialSupervisor, err := c.useCaseGetById.Execute(int32(supervisorID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Validación de pertenencia
	if initialSupervisor.User_id != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "not authorized to view this supervisor"})
		return
	}

	updateChan := make(chan domain.Supervisor)
	timeout := time.After(1 * time.Second)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			updatedSupervisor, err := c.useCaseGetById.Execute(int32(supervisorID))
			if err != nil {
				continue
			}
			if updatedSupervisor != initialSupervisor {
				updateChan <- updatedSupervisor
				return
			}
		}
	}()

	select {
	case updatedSupervisor := <-updateChan:
		ctx.JSON(http.StatusOK, gin.H{"supervisor": updatedSupervisor})
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"supervisor": initialSupervisor})
	}
}
