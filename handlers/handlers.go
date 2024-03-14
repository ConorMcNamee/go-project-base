package handlers

import (
	"fmt"
	"gobaseproject/database"
	"gobaseproject/models"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	// Restrict this route to POST only

	r.ParseForm()
	// Check to see if everything is of valid length's and types

	// Hash Password

	// Store in Database
	var u models.User
	u.FirstName = r.FormValue("first_name")
	u.LastName = r.FormValue("first_name")
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password") // Stores password in plain text. Must be changed. THIS IS PURELY FOR TESTING PURPOSES
	u.Admin = false
	u.Age = 18
	u.Joined_at = time.Now()
	u.Created_at = time.Now()
	u.Updated_at = time.Now()

	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}

	fmt.Println(u, r.FormValue("first_name"))
	fmt.Println("Creating User!")
	result, err := database.CreateNewUser(u)

	if err != nil {
		fmt.Println("Creating user failed")
	}

	if result == 0 {
		fmt.Println("Creating User was Successful!")
	}

	// Add JSON response
}
