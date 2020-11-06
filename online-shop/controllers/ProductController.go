package controllers

import (
	"../models"
	"encoding/json"
	"log"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	p := models.NewProduct(1, "laptop lenovo", 4000.43)
	w.Header().Set("Content-Type","application/json")
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}


