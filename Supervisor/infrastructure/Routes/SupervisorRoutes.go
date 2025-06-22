package routes

import (
	dependences "github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure/Dependences"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/supervisor")
	SaveSupervisor := dependences.GetSaveSupervisorController().Execute
	routes.POST("/", SaveSupervisor)
}
