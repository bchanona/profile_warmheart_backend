package routes

import (
	dependences "github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure/Dependences"
	"github.com/bchanona/profile_warmheart_backend/middlewares"
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
	LoginSupevisor := dependences.GetLoginController().Execute
	routes.POST("/", middlewares.AuthMiddleware(), SaveSupervisor)                         //user_Id
	routes.GET("/:id", middlewares.AuthMiddleware(), GetById)                              //se usa user_Id/ va a traer todo el json del supervisor pero los campos de contraseña y user_id estaran vacios
	routes.GET("/viewByUser", middlewares.AuthMiddleware(), GetSupervisorByUserID)         //user_Id
	routes.GET("/viewUser", middlewares.AuthSupervisorMiddleware(), GetUserBySupervisorID) //supervisor_Id
	routes.PUT("/", middlewares.AuthSupervisorMiddleware(), UpdateSupervisor)              //supervisor_Id
	routes.PUT("/changePassword", middlewares.AuthSupervisorMiddleware(), UpdatePassword)  //supervisor_Id
	routes.DELETE("/:id", middlewares.AuthMiddleware(), DeleteSupervisor)                  //user_Id
	routes.POST("/login", LoginSupevisor)
}
