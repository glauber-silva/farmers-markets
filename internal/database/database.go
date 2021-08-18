package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// create a string connection and return a database object
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Set up new database connection")

	username := os.Getenv("PSQL_USERNAME")
	passoword := os.Getenv("PSQL_PASSWORD")
	hostname := os.Getenv("PSQL_HOSTNAME")
	database := os.Getenv("PSQL_DATABASE")
	port := os.Getenv("PSQL_PORT")

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", hostname, port, username, database, passoword)

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err

	}

	return db, nil
}
