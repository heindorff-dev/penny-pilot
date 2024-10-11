package main

import (
	"api/database"
	"api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})

	route.RegisterRoutes(router)

	router.Use(gin.Logger())

	database.GetSession()

	router.Run(":3001")
}
