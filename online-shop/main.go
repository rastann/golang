package main

import (
	"./controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/product", controllers.GetProduct)


	http.ListenAndServe(":3000", r)
}
