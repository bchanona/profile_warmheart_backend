package main

import (
	userDependencies "github.com/bchanona/profile_warmheart_backend/User/infrastructure/dependencies"
	userRoutes "github.com/bchanona/profile_warmheart_backend/User/infrastructure/routes"
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

	userDependencies.Init()
	router := gin.Default()

	userRoutes.SetupUserRoutes(
		router,
		userDependencies.CreateUserController(),
		userDependencies.LoginUserController(),
		userDependencies.GetAllUsersController(),
		userDependencies.GetUserByIDController(),
		userDependencies.UpdateStatusController(),
		userDependencies.DeleteUserController(),
		userDependencies.GetJWTKey(),
	)

	router.Run(":8081")
}
