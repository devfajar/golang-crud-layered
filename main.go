package main

import (
	"fmt"
	"net/http"

	"github.com/devfajar/golang-rest-standlib/config"
	"github.com/devfajar/golang-rest-standlib/handler"
	"github.com/devfajar/golang-rest-standlib/helper"
	"github.com/devfajar/golang-rest-standlib/repository"
	"github.com/devfajar/golang-rest-standlib/route"
	"github.com/devfajar/golang-rest-standlib/service"

)

func main()  {
	fmt.Printf("Start server")

	// Initialize Database
	db := config.DatabaseConnection()


	// Repository Initialization
	bookRepository := repository.NewBookRepository(db)

	// Service Initialization
	bookService := service.NewBookService(bookRepository)

	// Handler Initialization
	bookHandler := handler.NewBookHandler(bookService)

	routes := route.NewRouter(bookHandler)

	routes.HandleFunc("/hello", handler.HelloHandler).Methods("GET")

	server := http.Server{
		Addr: "localhost:8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}