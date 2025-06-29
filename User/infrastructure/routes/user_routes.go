package routes

import (
    "github.com/bchanona/profile_warmheart_backend/User/infrastructure/controllers"
    "github.com/bchanona/profile_warmheart_backend/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, createUserCtrl *controllers.SaveUserController, 
    loginUserCtrl *controllers.LoginUserController, getAllUsersCtrl *controllers.GetAllUsersController,
    getByIDUserCtrl *controllers.GetByIDUserController, updateStatusCtrl *controllers.UpdateStatusController,
    deleteUserCtrl *controllers.DeleteUserController, jwtKey []byte) {

    userGroup := router.Group("/user")
    {
        userGroup.POST("/", createUserCtrl.Save)
        userGroup.POST("/login", loginUserCtrl.Login)
        
        userGroup.Use(middlewares.AuthMiddleware())
        {
            userGroup.GET("/", getAllUsersCtrl.GetAll)
            userGroup.PUT("/updateStatus", updateStatusCtrl.UpdateStatus)
            userGroup.DELETE("/:id", deleteUserCtrl.Delete)
        }
    }
}