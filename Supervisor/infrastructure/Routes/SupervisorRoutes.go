package routes

import (
	dependences "github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure/Dependences"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/supervisor")
	SaveSupervisor := dependences.GetSaveSupervisorController().Execute
	GetById := dependences.GetGetByIDController().Execute
	GetSupervisorByUserID := dependences.GetGetSupervisorByUserByIDController().Execute
	GetUserBySupervisorID := dependences.GetGetUserBySupervisorIDController().Execute
	UpdateSupervisor := dependences.GetUpdateSupervisorController().Execute
	UpdatePassword := dependences.GetUpdatePasswordController().Execute
	DeleteSupervisor := dependences.GetDeleteSupervisorController().Execute
	routes.POST("/", SaveSupervisor)
	routes.GET("/:id", GetById) //va a traer todo el json del supervisor pero los campos de contraseña y user_id estaran vacios
	routes.GET("/viewByUser/:id", GetSupervisorByUserID)
	routes.GET("/viewUser/:id", GetUserBySupervisorID)
	routes.PUT("/:id", UpdateSupervisor)
	routes.PUT("/changePassword/:id", UpdatePassword)
	routes.DELETE("/:id", DeleteSupervisor)
}
