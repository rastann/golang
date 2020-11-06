package main

import (
	"./routers"
	"net/http"
)

func main() {
	r := routers.IniRouter()
	http.ListenAndServe(":8080", r)
}
