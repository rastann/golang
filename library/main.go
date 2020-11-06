package main

import (
	"./driver"
	"./models"
	"./router"
	"database/sql"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

var books[] models.Book

var db *sql.DB

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	db = driver.ConnectDB()

	router := router.GetRouter(db)

	log.Println("Running server at port :8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
