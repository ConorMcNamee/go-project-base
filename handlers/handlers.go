package handlers

import (
	"fmt"
	"gobaseproject/database"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	user, err := database.GetUserByID(1)
	if err != nil {
		fmt.Println("Cannot find user with that ID")
	}

	fmt.Fprintf(w, "Hey: %s", user.FirstName)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

}
