package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/raynaaliyya/enigma-laundry/utils"
)

var (
	dbHost     = utils.DotEnv("DB_HOST")
	dbPort     = utils.DotEnv("DB_PORT")
	dbUser     = utils.DotEnv("DB_USER")
	dbPassword = utils.DotEnv("DB_PASSWORD")
	dbName     = utils.DotEnv("DB_NAME")
	sslMode    = utils.DotEnv("SSL_MODE")
)

var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)

func NewDatabase() *sql.DB {

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Connection failed %v", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully connected")
	}

	return db
}
