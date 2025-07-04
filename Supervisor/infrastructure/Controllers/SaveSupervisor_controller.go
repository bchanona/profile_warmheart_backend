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
	userIDRaw, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}
	userID := userIDRaw.(int)

	var req domain.CreateSupervisorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	supervisor := domain.Supervisor{
		Name:     req.Name,
		Surnames: req.Surnames,
		Email:    req.Email,
		Password: req.Password,
		User_id:  userID,
	}

	err := c.CreateUseCase.Execute(supervisor)
	if err != nil {
		switch err {
		case domain.ErrSupervisorAlreadyExists:
			ctx.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		case domain.ErrInvalidInput:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing fields"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not register supervisor"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "supervisor registered successfully"})
}
