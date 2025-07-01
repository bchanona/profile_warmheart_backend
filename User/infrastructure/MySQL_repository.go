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