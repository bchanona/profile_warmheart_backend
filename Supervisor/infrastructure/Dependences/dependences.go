package dependences

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bchanona/profile_warmheart_backend/Supervisor/application"
	"github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure"
	controllers "github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure/Controllers"
	"github.com/bchanona/profile_warmheart_backend/helpers"
)

var (
	mySQL infrastructure.MySQLRepository
	db    *sql.DB
)

func Init() {
	db, err := helpers.ConnectMySQL()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	mySQL = *infrastructure.NewMySQLRepository(db)
}
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}
func GetSaveSupervisorController() *controllers.SaveUserController {
	caseSaveSupervisor := application.SaveSupervisor(&mySQL)
	return controllers.NewSaveUserController(caseSaveSupervisor)
}
func GetGetByIDController() *controllers.GetByIdController {
	caseGetByID := application.GetByID(&mySQL)
	return controllers.NewGetByIdProductController(caseGetByID)
}
func GetGetSupervisorByUserByIDController() *controllers.GetSupervisorByUserIdController {
	caseGetSupervisorByUserID := application.NewGetSupervisorByUserIDUseCase(&mySQL)
	return controllers.NewGetSupervisorByUserIdController(caseGetSupervisorByUserID)
}
func GetGetUserBySupervisorIDController() *controllers.GetUserBySupervisorIDController {
	caseGetUserBySupervisorID := application.NewGetUserBySupervisorIDUseCase(&mySQL)
	return controllers.NewGetUserBySupervisorIDController(caseGetUserBySupervisorID)
}
func GetUpdateSupervisorController() *controllers.UpdateSupervisorController {
	caseUpdateSupervisor := application.NewUpdateSupervisorUseCase(&mySQL)
	return controllers.NewUpdateSupervisorController(caseUpdateSupervisor)
}
func GetUpdatePasswordController() *controllers.UpdatePasswordController {
	caseUpdatePassword := application.NewUpdateSupervisorPasswordUseCase(&mySQL)
	return controllers.NewUpdatePasswordController(caseUpdatePassword)
}
func GetDeleteSupervisorController() *controllers.DeleteSupervisorController {
	caseDeleteSupervisor := application.NewDeleteSupervisorUseCase(&mySQL)
	return controllers.NewDeleteSupervisorController(caseDeleteSupervisor)
}
