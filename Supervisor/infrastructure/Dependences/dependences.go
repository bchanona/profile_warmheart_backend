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
