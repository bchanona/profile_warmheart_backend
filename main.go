package main

import (
	userDependencies "github.com/bchanona/profile_warmheart_backend/User/infrastructure/dependencies"
	userRoutes "github.com/bchanona/profile_warmheart_backend/User/infrastructure/routes"
	"github.com/bchanona/profile_warmheart_backend/helpers"
	supervisor "github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure/Dependences"
	routeSupervisor "github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure/Routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := helpers.ConnectMySQL()
	if db == nil {
		panic("Error connecting to the database")
	} else {
		println("Successful connection to the database")
	}

	userDependencies.Init()
  supervisor.Init()
  defer supervisor.CloseDB()
  


	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	routeSupervisor.Routes(r)
  userRoutes.SetupUserRoutes(
		r,
		userDependencies.CreateUserController(),
		userDependencies.LoginUserController(),
		userDependencies.GetAllUsersController(),
		userDependencies.GetUserByIDController(),
		userDependencies.UpdateStatusController(),
		userDependencies.DeleteUserController(),
		userDependencies.GetJWTKey(),
	)
	r.Run(":8081")

}
