package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) {
	if sessions.Default(c).Get("profile") == nil {
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		c.Next()
	}
}
