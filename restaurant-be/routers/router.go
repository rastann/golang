package routers

import (
	"github.com/gorilla/mux"
	dishHandler "./handlers/dishes"
)

func IniRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/dishes", dishHandler.GetDishes).Methods("GET")
	return r
}
