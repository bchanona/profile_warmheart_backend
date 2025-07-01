package main

import (
	"github.com/bchanona/profile_warmheart_backend/User/infrastructure/dependencies"
	"github.com/bchanona/profile_warmheart_backend/User/infrastructure/routes"
	"github.com/bchanona/profile_warmheart_backend/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := helpers.ConnectMySQL()
	if db == nil {
		panic("Error connecting to the database")
	} else {
		println("Successful connection to the database")
	}

	dependencies.Init()
	router := gin.Default()

	routes.SetupUserRoutes(
		router,
		dependencies.CreateUserController(),
		dependencies.LoginUserController(),
		dependencies.GetAllUsersController(),
		dependencies.GetUserByIDController(),
		dependencies.UpdateStatusController(),
		dependencies.DeleteUserController(),
		dependencies.GetJWTKey(),
	)

	router.Run(":8000")
}
