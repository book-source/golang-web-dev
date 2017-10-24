package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8081", http.FileServer(http.Dir("./starting-files")))
	if err != nil {
		log.Fatal(err)
	}
}
