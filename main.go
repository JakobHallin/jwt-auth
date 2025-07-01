package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting up program")

	http.HandleFunc("/signup", SignupHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/protected", ProtectedHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
