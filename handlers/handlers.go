package handlers

import (
	"fmt"
	"gobaseproject/database"
	"gobaseproject/models"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Home(w http.ResponseWriter, r *http.Request) {
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Restrict this route to POST only
	r.ParseForm()
	// Check to see if everything is of valid length's and types

	// Hash Password
	password, err := hashPassword(r.FormValue("password"))
	if err != nil {
		return
	}

	// Store in Database
	var u models.User
	u.FirstName = r.FormValue("first_name")
	u.LastName = r.FormValue("first_name")
	u.Email = r.FormValue("email")
	u.Password = password // Stores password in plain text. Must be changed. THIS IS PURELY FOR TESTING PURPOSES
	u.Admin = false
	u.Age = 18
	u.Joined_at = time.Now()
	u.Created_at = time.Now()
	u.Updated_at = time.Now()

	fmt.Println("Creating User!")
	result, err := database.CreateNewUser(u)

	if err != nil {
		fmt.Println("Creating user failed")
		return
	}

	if result == 0 {
		fmt.Println("Creating User was Successful!")
	}

	// Add JSON response
}
