package dependencies

import (
	"log"
	"os"

	"github.com/bchanona/profile_warmheart_backend/User/application"
	"github.com/bchanona/profile_warmheart_backend/User/infrastructure"
	"github.com/bchanona/profile_warmheart_backend/User/infrastructure/controllers"
	"github.com/bchanona/profile_warmheart_backend/helpers"
)

var (
	createUserUseCase     *application.SaveUserUseCase
	loginUserUseCase      *application.LoginUserUseCase
	getAllUsersUseCase    *application.GetAllUsersUseCase
	getUserByIDUseCase    *application.GetByIDUserUseCase
	updateStatusUseCase   *application.UpdateStatusUseCase
	deleteUserUseCase     *application.DeleteUserUseCase
	jwtKey                []byte
)

func Init() {
	db, err := helpers.ConnectMySQL()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	userRepo := infrastructure.NewMySQLRepository(db)

	createUserUseCase = application.NewSaveUserUseCase(userRepo)
	loginUserUseCase = application.NewLoginUserUseCase(userRepo)
	getAllUsersUseCase = application.NewGetAllUsersUseCase(userRepo)
	getUserByIDUseCase = application.NewGetByIDUserUseCase(userRepo)
	updateStatusUseCase = application.NewUpdateStatusUseCase(userRepo)
	deleteUserUseCase = application.NewDeleteUserUseCase(userRepo)

	// Load JWT key
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY environment variable not set")
	}
}

func CreateUserController() *controllers.SaveUserController {
	return controllers.NewSaveUserController(createUserUseCase)
}

func LoginUserController() *controllers.LoginUserController {
	return controllers.NewLoginUserController(loginUserUseCase, jwtKey)
}

func GetAllUsersController() *controllers.GetAllUsersController {
	return controllers.NewGetAllUsersController(getAllUsersUseCase)
}

func GetUserByIDController() *controllers.GetByIDUserController {
	return controllers.NewGetByIDUserController(getUserByIDUseCase)
}

func UpdateStatusController() *controllers.UpdateStatusController {
	return controllers.NewUpdateStatusController(updateStatusUseCase)
}

func DeleteUserController() *controllers.DeleteUserController {
	return controllers.NewDeleteUserController(deleteUserUseCase)
}

func GetJWTKey() []byte {
	return jwtKey
}