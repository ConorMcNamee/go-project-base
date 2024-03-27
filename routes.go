package main

import (
	"gobaseproject/handlers"
	"gobaseproject/middleware"
	"net/http"
)

func HandleRoutes() {

	http.HandleFunc("/", handlers.Home)

	loginHandler := http.HandlerFunc(handlers.Login)
	http.Handle("/login", middleware.OnlyPost(loginHandler))

	http.HandleFunc("/logout", handlers.Logout)

	registerHandler := http.HandlerFunc(handlers.Register)
	http.Handle("/register", middleware.OnlyPost(registerHandler))
}
