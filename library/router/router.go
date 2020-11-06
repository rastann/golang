package router

import (
	"../controllers"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter(db *sql.DB) http.Handler {
	router := mux.NewRouter()

	controller := new(controllers.Controller)
	controller.Init()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetOneBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	return router
}
