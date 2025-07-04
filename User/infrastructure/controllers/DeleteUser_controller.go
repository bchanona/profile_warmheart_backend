package controllers

import (
	"net/http"
	"strconv"

	"github.com/bchanona/profile_warmheart_backend/User/application"
	"github.com/bchanona/profile_warmheart_backend/User/domain"
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
    UseCase *application.DeleteUserUseCase
}

func NewDeleteUserController(useCase *application.DeleteUserUseCase) *DeleteUserController {
    return &DeleteUserController{UseCase: useCase}
}

func (ctrl *DeleteUserController) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := ctrl.UseCase.Delete(id); err != nil {
        if err == domain.ErrUserNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}