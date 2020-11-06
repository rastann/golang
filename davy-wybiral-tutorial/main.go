package main

import (
	"./models"
	"./routes"
	"./view"
	"context"
	"net/http"
)

var ctx = context.Background()

func main() {
	models.Init()
	view.LoadTemplates("templates/*.html")
	r := routes.NewRouter(&ctx)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}