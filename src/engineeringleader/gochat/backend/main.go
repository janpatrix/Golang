package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Print("Server started at localhost:4444")
	if err := http.ListenAndServe(":4444", nil); err != nil {
		log.Fatal(err)
	}
}
