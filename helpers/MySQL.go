package helpers

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectMySQL() (db *sql.DB, err error) {
	loadVerify := godotenv.Load()

	if loadVerify != nil {
		fmt.Println("Error loading .env file")
	}

	user_DB := os.Getenv("DB_USER")
	password_DB := os.Getenv("DB_PASSWORD")
	host_DB := os.Getenv("DB_HOST")
	name_DB := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user_DB, password_DB, host_DB, name_DB)

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Error opening connection: %s\n", err.Error())
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		fmt.Printf("Error verifying connection: %s\n", err.Error())
		return nil, err
	}

	return db, nil

}
