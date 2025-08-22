package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/devfajar/golang-rest-standlib/helper"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)


func DatabaseConnection() *sql.DB {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)


	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Database connection established")
	return db
}