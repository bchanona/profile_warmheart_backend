package controllers

import (
	"net/http"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/gin-gonic/gin"
)

type SaveUserController struct {
	CreateUseCase *application.SaveSupervisorUseCase
}

func NewSaveUserController(CreateuseCase *application.SaveSupervisorUseCase) *SaveUserController {
	return &SaveUserController{CreateUseCase: CreateuseCase}
}

func (c *SaveUserController) Execute(ctx *gin.Context) {
	var supervisor domain.Supervisor
	if err := ctx.ShouldBindJSON(&supervisor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	err := c.CreateUseCase.Execute(supervisor)
	if err != nil {
		if err == domain.ErrSupervisorAlreadyExists {
			ctx.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not register user"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}
