package router

import (
	"api/rest/auth"
	"api/rest/handlers/auth/callback"
	"api/rest/handlers/auth/login"
	"api/rest/handlers/auth/logout"
	"api/rest/handlers/auth/user"
	"api/rest/handlers/expense"
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func New(a *auth.Authenticator) *gin.Engine {
	r := gin.New()

	//Middleware
	r.Use(gin.Logger())
	r.RemoveExtraSlash = true

	//Auth0
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth-session", store))

	//r.Static("/public", "web/static")
	//r.LoadHTMLGlob("web/template/*")

	r.GET("api/v1/login", login.Handler(a))
	r.GET("api/v1/callback", callback.Handler(a))
	r.GET("api/v1/user", user.Handler)
	r.GET("api/v1/logout", logout.Handler)

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
