package main

import (
	"api/internal/database"
	router "api/rest"
	"api/rest/auth"
	"log"
	"net/http"
)

func main() {
	database.Init()
	cassandraSession := database.Session
	defer cassandraSession.Close()

	auth, err := auth.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	r := router.New(auth)
	log.Print("Server listening on http://localhost:3001/")
	if err := http.ListenAndServe("0.0.0.0:3001", r); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
