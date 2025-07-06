package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
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
	existingUser, err := r.GetSupervisorByEmail(supervisor.Email)
	if err == nil && existingUser.User_id != 0 {
		return domain.ErrSupervisorAlreadyExists
	}

	// Encripta contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(supervisor.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO SUPERVISORS
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

func (r *MySQLRepository) Delete(id int) error {
	return errors.New("not implemented yet")
}
func (mysql *MySQLRepository) GetByIDSupervisor(id int32) (domain.Supervisor, error) {
	var SupervisorById domain.Supervisor

	query := "SELECT supervisor_id, name, surnames,email FROM SUPERVISORS WHERE supervisor_id=?"
	row := mysql.db.QueryRow(query, id)

	err := row.Scan(&SupervisorById.Supervisor_id, &SupervisorById.Name, &SupervisorById.Surnames, &SupervisorById.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return SupervisorById, fmt.Errorf("Supervisor no encontrado con id: ", id)
		}
		return SupervisorById, err
	}

	return SupervisorById, nil
}
func (r *MySQLRepository) GetSupervisorsByUserID(userID int32) ([]domain.GetSupervisorByUserID, error) {
	query := `SELECT user_id, supervisor_id, name, surnames, email FROM SUPERVISORS WHERE user_id = ?`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var supervisors []domain.GetSupervisorByUserID

	for rows.Next() {
		var sup domain.GetSupervisorByUserID
		err := rows.Scan(&sup.User_id, &sup.Supervisor_id, &sup.Name, &sup.Surnames, &sup.Email)
		if err != nil {
			return nil, err
		}
		supervisors = append(supervisors, sup)
	}

	return supervisors, nil
}
func (r *MySQLRepository) GetUserBySupervisorID(id int32) (domain.GetUserBySupervisorID, error) {
	query := `
		SELECT 
			s.supervisor_id, u.user_id, u.device_id,u.name,u.surnames,u.email,u.premium FROM SUPERVISORS s 
		INNER JOIN USERS u ON u.user_id = s.user_id 
		WHERE s.supervisor_id = ? 
		LIMIT 1`

	var user domain.GetUserBySupervisorID
	err := r.db.QueryRow(query, id).Scan(
		&user.Supervisor_id,
		&user.User_id,
		&user.Device_id,
		&user.Name,
		&user.Surnames,
		&user.Email,
		&user.Premium,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, domain.ErrUserNotFound
		}
		return user, err
	}

	return user, nil
}
func (r *MySQLRepository) UpdateSupervisor(id int, data domain.UpdateSupervisor) error {
	query := `UPDATE SUPERVISORS SET name = ?, surnames = ?, email = ? WHERE supervisor_id = ?`

	result, err := r.db.Exec(query, data.Name, data.Surnames, data.Email, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrSupervisorNotFound
	}

	return nil
}
func (r *MySQLRepository) GetSupervisorByEmail(email string) (domain.Supervisor, error) {
	query := `SELECT supervisor_id, name, surnames, email FROM SUPERVISORS WHERE email = ? LIMIT 1`

	var s domain.Supervisor
	err := r.db.QueryRow(query, email).Scan(
		&s.Supervisor_id,
		&s.Name,
		&s.Surnames,
		&s.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s, domain.ErrSupervisorNotFound
		}
		return s, err
	}
	return s, nil
}
func (r *MySQLRepository) UpdateSupervisorPassword(id int, data domain.UpdatePassword) error {
	if data.Password == "" {
		return domain.ErrInvalidInput
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `UPDATE SUPERVISORS SET password = ? WHERE supervisor_id = ?`
	result, err := r.db.Exec(query, string(hashedPassword), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrSupervisorNotFound
	}

	return nil
}
func (r *MySQLRepository) DeleteSupervisor(supervisorID int, userID int) error {
	query := `DELETE FROM SUPERVISORS WHERE supervisor_id = ? AND user_id = ?`

	result, err := r.db.Exec(query, supervisorID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrSupervisorNotFound
	}

	return nil
}
func (sql *MySQLRepository) LoginSupervisors(email string, password string) (domain.Supervisor, error) {
	var supervisor domain.Supervisor
	var hashedPassword string

	query := "SELECT supervisor_id, name, surnames, email, password, user_id FROM SUPERVISORS WHERE email = ?"

	err := sql.db.QueryRow(query, email).Scan(&supervisor.Supervisor_id, &supervisor.Name, &supervisor.Surnames, &supervisor.Email, &hashedPassword, &supervisor.User_id)
	if err != nil {
		return domain.Supervisor{}, fmt.Errorf("error al buscar supervisor: %w", err)

	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return domain.Supervisor{}, errors.New("contraseña incorrecta")
	}

	supervisor.Password = ""
	return supervisor, nil
}
