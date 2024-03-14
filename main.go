package main

import (
	"fmt"
	"gobaseproject/config"
	"gobaseproject/database"
	"gobaseproject/driver"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const portNumber = ":8080"

var app config.AppConfig

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	// var app config.AppConfig
	HandleRoutes()
	fmt.Println("Connecting to database...")
	dsn := fmt.Sprintf("host=localhost port=5432 dbname=%s user=%s password=%s", goDotEnvVariable("POSTGRES_DB"), goDotEnvVariable("POSTGRES_USER"), goDotEnvVariable("POSTGRES_PASSWORD"))
	db, err := driver.ConnectSQL(dsn)

	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully!")

	app.DB = db
	database.SetDB(&app)

	fmt.Println("Listening on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
