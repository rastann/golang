package dishes

import (
	"io"
	"net/http"
)

func GetDishes(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello")
}