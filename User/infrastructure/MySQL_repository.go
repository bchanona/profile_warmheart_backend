// infrastructure/mysql_repository.go
package infrastructure

import (
	"database/sql"
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
    // Verifica si el usuario ya existe
    existingUser, err := r.GetByEmail(user.Email)
    if err == nil && existingUser.User_id != 0 {
        return domain.ErrUserAlreadyExists
    }

    // Encripta contraseña
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    query := `INSERT INTO usuario 
        (nombre, apellidos, correo, password, premium, id_dispositivo) 
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
    query := `SELECT user_id, nombre, apellidos, correo, password, premium, id_dispositivo 
              FROM usuario WHERE correo = ?`
    
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
    query := `SELECT user_id, nombre, apellidos, correo, password, premium, id_dispositivo
              FROM usuario WHERE user_id = ?`

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


func (r *MySQLRepository) Update(user domain.User) error {
    return nil 
}

func (r *MySQLRepository) UpdatePassword(id int, newPassword string) error {
    return nil 
}

func (r *MySQLRepository) UpdateStatus(id int, premium bool) error {
    return nil 
}

func (r *MySQLRepository) Delete(id int) error {
    return nil
}
