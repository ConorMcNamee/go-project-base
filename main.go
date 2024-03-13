package main

import (
	"fmt"
	"gobaseproject/config"
	"gobaseproject/database"
	"gobaseproject/driver"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	// var app config.AppConfig
	HandleRoutes()
	fmt.Println("Connecting to database...")

	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=mydatabase user=myuser password=mypassword")

	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully!")

	app.DB = db
	database.SetDB(&app)

	fmt.Println("Listening on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
