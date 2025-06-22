package infrastructure

import (
	"database/sql"
	"errors"
	"log"

	domain "github.com/bchanona/profile_warmheart_backend/Supervisor/Domain"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) Save(supervisor domain.Supervisor) error {
	// Verifica si el supervisor ya existe
	existingUser, err := r.GetByEmail(supervisor.Email)
	if err == nil && existingUser.User_id != 0 {
		return domain.ErrSupervisorAlreadyExists
	}

	// Encripta contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(supervisor.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO supervisors
		(name, surnames, email, password, user_id) 
		VALUES (?, ?, ?, ?, ?)`

	_, err = r.db.Exec(query,
		supervisor.Name,
		supervisor.Surnames,
		supervisor.Email,
		string(hashedPassword),
		supervisor.User_id,
	)

	if err != nil {
		// Validar si el error fue por email duplicado (clave única)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			// 1062 = Duplicate entry
			return domain.ErrSupervisorAlreadyExists
		}
		log.Println("Error saving supervisor:", err)
		return err
	}

	return nil
}
func (r *MySQLRepository) GetByEmail(email string) (domain.Supervisor, error) {
	var supervisor domain.Supervisor

	query := `SELECT id, name, surnames, email, password, user_id 
	          FROM SUPERVISORS WHERE email = ? LIMIT 1`

	row := r.db.QueryRow(query, email)
	err := row.Scan(
		&supervisor.Supervisor_id,
		&supervisor.Name,
		&supervisor.Surnames,
		&supervisor.Email,
		&supervisor.Password,
		&supervisor.User_id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Email no existe — válido para hacer inserción
			return domain.Supervisor{}, sql.ErrNoRows
		}
		// Error inesperado
		return domain.Supervisor{}, err
	}

	return supervisor, nil
}
func (r *MySQLRepository) Delete(id int) error {
	return errors.New("not implemented yet")
}
