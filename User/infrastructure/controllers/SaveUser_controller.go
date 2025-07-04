package controllers

import (
	"net/http"

	"github.com/bchanona/profile_warmheart_backend/User/application"
	"github.com/bchanona/profile_warmheart_backend/User/domain"
	"github.com/gin-gonic/gin"
)

type SaveUserController struct {
	UseCase *application.SaveUserUseCase
}

func NewSaveUserController(useCase *application.SaveUserUseCase) *SaveUserController {
    return &SaveUserController{UseCase: useCase}
}

func (c *SaveUserController) Save(ctx *gin.Context) {
    var user domain.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
        return
    }

    err := c.UseCase.Save(user)
    if err != nil {
        if err == domain.ErrUserAlreadyExists {
            ctx.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not register user"})
        }
        return
    }
    
    ctx.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}
