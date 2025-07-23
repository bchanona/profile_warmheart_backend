// infrastructure/mysql_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bchanona/profile_warmheart_backend/User/domain"
	"golang.org/x/crypto/bcrypt"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) Save(user domain.User) error {
	existingUser, err := r.GetByEmail(user.Email)
	if err == nil && existingUser.User_id != 0 {
		return domain.ErrUserAlreadyExists
	}

	// Encripta contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO USERS 
        (name, surnames, email, password, premium, device_id) 
        VALUES (?, ?, ?, ?, ?, ?)`

	_, err = r.db.Exec(query,
		user.Name,
		user.Surnames,
		user.Email,
		string(hashedPassword),
		user.Premium,
		user.Device_id)

	if err != nil {
		log.Println("Error saving user:", err)
		return err
	}
	return nil
}

func (r *MySQLRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	query := `SELECT user_id, name, surnames, email, password, premium, device_id 
              FROM USERS WHERE email = ?`

	err := r.db.QueryRow(query, email).Scan(
		&user.User_id,
		&user.Name,
		&user.Surnames,
		&user.Email,
		&user.Password,
		&user.Premium,
		&user.Device_id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

func (r *MySQLRepository) GetByID(id int) (domain.User, error) {
	var user domain.User
	query := `SELECT user_id, name, surnames, email, password, premium, device_id
              FROM USERS WHERE user_id = ?`

	err := r.db.QueryRow(query, id).Scan(
		&user.User_id,
		&user.Name,
		&user.Surnames,
		&user.Email,
		&user.Password,
		&user.Premium,
		&user.Device_id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *MySQLRepository) GetAll() ([]domain.User, error) {
	query := `SELECT user_id, name, surnames, email, premium, device_id FROM USERS`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.User_id,
			&user.Name,
			&user.Surnames,
			&user.Email,
			&user.Premium,
			&user.Device_id,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *MySQLRepository) UpdateStatus(id int, premium bool) error {
	query := `UPDATE USERS SET premium = ? WHERE user_id = ?`
	_, err := r.db.Exec(query, premium, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrUserNotFound
		}
		return err
	}
	return nil
}

func (r *MySQLRepository) Delete(id int) error {
	query := `DELETE FROM USERS WHERE user_id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrUserNotFound
		}
		return err
	}
	return nil
}

func (r *MySQLRepository) SaveNotification(notification domain.SaveNotifications) error {
	query := `INSERT INTO NOTIFICATIONS (user_id,body,reading) VALUES (?,?,?)`
	_, err := r.db.Exec(query, notification.User_id, notification.BodyMessage, notification.Reading)
	if err != nil {
		log.Println("Error saving user:", err)
		return err
	}
	return nil

}
func (r *MySQLRepository) GetNotification(user_id int) ([]domain.GetNotifications, error) {
	var notifications []domain.GetNotifications

	query := `SELECT user_id,body,reading,date,time FROM NOTIFICATIONS WHERE user_id = ?`

	rows, err := r.db.Query(query, user_id)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var notification domain.GetNotifications

		err := rows.Scan(
			&notification.User_id,
			&notification.BodyMessage,
			&notification.Reading,
			&notification.Date,
			&notification.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}
