package dependencies

import (
	"log"

	"github.com/bchanona/profile_warmheart_backend/User/application"
	"github.com/bchanona/profile_warmheart_backend/User/infrastructure"
	"github.com/bchanona/profile_warmheart_backend/User/infrastructure/controllers"
	"github.com/bchanona/profile_warmheart_backend/helpers"
)

var (
	userService *application.SaveUserUseCase
)

func Init() {
	db, err := helpers.ConnectMySQL()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	userRepo := infrastructure.NewMySQLRepository(db)
	userService = application.NewSaveUserUseCase(userRepo)
}

func UserController() *controllers.SaveUserController {
	return controllers.NewSaveUserController(userService)
}