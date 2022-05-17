package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	host     = loadDotEnvVariable("POSTGRES_HOST")
	port     = loadDotEnvVariable("POSTGRES_PORT")
	user     = loadDotEnvVariable("POSTGRES_USER")
	password = loadDotEnvVariable("POSTGRES_PASSWORD")
	dbname   = loadDotEnvVariable("POSTGRES_DB")
)

func loadDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

func DBConnect() {
	// Connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

}
