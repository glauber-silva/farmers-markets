package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
/*
	TODO: Implement Repository to change between databases, mainly to separate the test database (sqlite)
 */

// create a string connection and return a database object
func NewDatabase() (*gorm.DB, error) {
	log.Info("Set up new database connection")

	username := os.Getenv("PSQL_USERNAME")
	passoword := os.Getenv("PSQL_PASSWORD")
	hostname := os.Getenv("PSQL_HOSTNAME")
	database := os.Getenv("PSQL_DATABASE")
	port := os.Getenv("PSQL_PORT")


	connection := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=%s sslmode=disable", hostname,
		username, passoword, database, port)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
