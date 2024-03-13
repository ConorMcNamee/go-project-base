package main

import (
	"gobaseproject/handlers"
	"net/http"
)

func HandleRoutes() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/register", handlers.Register)
}
