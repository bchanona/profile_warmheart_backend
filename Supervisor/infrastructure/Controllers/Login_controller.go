package controllers

import (
	"net/http"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/bchanona/profile_warmheart_backend/helpers"
	"github.com/gin-gonic/gin"
)

type LogInController struct {
	useCase *application.LoginUseCase
}

func NewLogInController(useCase *application.LoginUseCase) *LogInController {
	return &LogInController{useCase: useCase}
}

func (controller *LogInController) Execute(ctx *gin.Context) {
	var supervisor domain.Supervisor
	if err := ctx.ShouldBindJSON(&supervisor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	email := supervisor.Email
	password := supervisor.Password

	authenticatedSupervisor, err := controller.useCase.Execute(email, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	if authenticatedSupervisor.Supervisor_id == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ID de supervisor"})
		return
	}

	token, err := helpers.GenerateSupervisorJWT(authenticatedSupervisor.Supervisor_id, authenticatedSupervisor.User_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"supervisor": authenticatedSupervisor, "token": token})
}
