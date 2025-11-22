package routes

import (
	"github.com/bchanona/profile_warmheart_backend/User/infrastructure/controllers"
	"github.com/bchanona/profile_warmheart_backend/middlewares"
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
)

func SetupUserRoutes(router *gin.Engine, createUserCtrl *controllers.SaveUserController,
	loginUserCtrl *controllers.LoginUserController, getAllUsersCtrl *controllers.GetAllUsersController,
	getByIDUserCtrl *controllers.GetByIDUserController, updateStatusCtrl *controllers.UpdateStatusController,
	deleteUserCtrl *controllers.DeleteUserController, saveNotification *controllers.SaveNotificationController,
	getNotification *controllers.GetNotificationController, jwtKey []byte) {

	router.GET("/", func(c *gin.Context) {
    serverName := os.Getenv("SERVER_NAME")
    if serverName == "" {
        serverName = "NombreServidorNoDefinido"
    }

    c.JSON(200, gin.H{
        "message": fmt.Sprintf("Ejecutando servidor %s", serverName),
    })
})

	userGroup := router.Group("/user")
	{
		userGroup.POST("/", createUserCtrl.Save)
		userGroup.POST("/login", loginUserCtrl.Login)
		userGroup.POST("/saveNotification", saveNotification.Execute)
		userGroup.GET("/getNotifications", middlewares.AuthSupervisorMiddleware(), getNotification.Execute) //hola toy aqui :)
		userGroup.Use(middlewares.AuthMiddleware())
		{
			userGroup.GET("/", getAllUsersCtrl.GetAll)
			userGroup.PUT("/updateStatus", updateStatusCtrl.UpdateStatus)
			userGroup.DELETE("/:id", deleteUserCtrl.Delete)
		}
	}
}
