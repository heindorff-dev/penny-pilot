package main

import (
	"api/internal/database"
	"api/internal/router"
)

func main() {
	database.MustInit()
	cassandraSession := database.Session
	defer cassandraSession.Close()
	r := router.New()
	r.Run(":3001")
}
