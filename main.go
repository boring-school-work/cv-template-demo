package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":3000", http.FileServer(http.Dir("./")))
	if err != nil {
		log.Fatal(err)
	}
}
