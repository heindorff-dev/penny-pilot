package router

import (
	"api/internal/resource/expense"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	//Middleware
	r.Use(gin.Logger())
	r.RemoveExtraSlash = true

	//Routes
	r.GET("api/v1/ping", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	//Expenses
	expenses := r.Group("api/v1/expenses")
	expenses.GET("/:id", func(ctx *gin.Context) {
		expense.GetExpense(ctx)
	})
	expenses.GET("/", func(ctx *gin.Context) {
		expense.GetExpenses(ctx)
	})

	return r
}
