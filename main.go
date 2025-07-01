package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"
	_ "github.com/lib/pq" //driver
)

func main() {
	log.Println("Starting up program")
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("FAILED to open database", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}
	store := &UserStore{db: db}
	server := &Server{store: store}
	if err := store.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}
	http.HandleFunc("/signup", server.SignupHandler)
	http.HandleFunc("/login", server.LoginHandler)
	http.HandleFunc("/protected", server.ProtectedHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
