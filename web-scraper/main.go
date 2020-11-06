package main

import (
	"log"
	"net/http"
)

func main() {
	url := "https://www.flashscore.ro/"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)

}