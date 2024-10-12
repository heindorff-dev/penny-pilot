package router

import (
	"api/pkg/handler/expense"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

func New(j *jwtmiddleware.JWTMiddleware) *gin.Engine {
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
	expenses.Use(adapter.Wrap(j.CheckJWT))
	expenses.GET("/:id", func(c *gin.Context) {
		expense.GetExpense(c)
	})
	expenses.GET("/", func(c *gin.Context) {
		expense.GetExpenses(c)
	})

	return r
}
