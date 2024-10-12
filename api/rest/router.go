package router

import (
	"api/rest/handlers/expense"

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
	expenses.GET("/:id", func(c *gin.Context) {
		expense.GetExpense(c)
	})
	expenses.GET("/", func(c *gin.Context) {
		expense.GetExpenses(c)
	})

	return r
}
