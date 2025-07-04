package controllers

import (
    "net/http"
    "github.com/bchanona/profile_warmheart_backend/User/application"
    "github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
    UseCase *application.GetAllUsersUseCase
}

func NewGetAllUsersController(useCase *application.GetAllUsersUseCase) *GetAllUsersController {
    return &GetAllUsersController{UseCase: useCase}
}

func (ctrl *GetAllUsersController) GetAll(c *gin.Context) {
    users, err := ctrl.UseCase.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}