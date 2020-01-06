package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// CreatePostgresConnection establishes a new postgres gorm connection.
func CreatePostgresConnection() (*gorm.DB, error) {
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		port     = os.Getenv("DB_PORT")
		dbname   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			host, port, user, dbname, password,
		),
	)
}
