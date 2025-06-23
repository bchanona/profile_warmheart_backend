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

func (getByIdSupervisor *GetByIdController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	//long polling
	initialSupervisor, err := getByIdSupervisor.useCaseGetById.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateChan := make(chan domain.Supervisor)
	timeout := time.After(1 * time.Second)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			updatedSupervisor, err := getByIdSupervisor.useCaseGetById.Execute(int32(id))
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
