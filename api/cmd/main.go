package main

import (
	"api/internal/database"
	router "api/rest"
)

func main() {
	database.Init()
	cassandraSession := database.Session
	defer cassandraSession.Close()
	r := router.New()
	r.Run(":3001")
}
