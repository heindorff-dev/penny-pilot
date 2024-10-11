package main

import (
	"api/internal/router"
)

func main() {
	r := router.New()
	r.Run(":3001")
}
