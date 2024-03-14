package main

import (
	"gobaseproject/handlers"
	"gobaseproject/middleware"
	"net/http"
)

func HandleRoutes() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logout", handlers.Logout)

	registerHandler := http.HandlerFunc(handlers.Register)
	http.Handle("/register", middleware.OnlyPost(registerHandler))
}
