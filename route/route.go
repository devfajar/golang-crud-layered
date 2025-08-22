package route

import (
	"github.com/devfajar/golang-rest-standlib/handler"
	"github.com/gorilla/mux"
)

func NewRouter(bookHandler *handler.BookHandler) *mux.Router {
	routes := mux.NewRouter()

	// Define your routes here
	routes.HandleFunc("/hello", handler.HelloHandler).Methods("GET")

	routes.HandleFunc("/books", bookHandler.Create).Methods("POST")
	routes.HandleFunc("/books/{bookId}", bookHandler.Update).Methods("PUT")
	routes.HandleFunc("/books/{bookId}", bookHandler.Delete).Methods("DELETE")
	routes.HandleFunc("/books", bookHandler.FindAll).Methods("GET")
	routes.HandleFunc("/books/{bookId}", bookHandler.FindById).Methods("GET")

	return routes
}