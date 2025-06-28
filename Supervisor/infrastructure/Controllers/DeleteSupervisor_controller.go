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
	supervisorIDParam := ctx.Param("id")
	supervisorID, err := strconv.Atoi(supervisorIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de supervisor inválido"})
		return
	}

	userIDRaw, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o no autorizado"})
		return
	}
	userID := userIDRaw.(int)

	// Ejecutar solo si el supervisor pertenece al usuario autenticado
	err = c.useCase.Execute(supervisorID, userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrSupervisorNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": "El supervisor no existe o no pertenece a este usuario"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el supervisor"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Supervisor eliminado correctamente"})
}
