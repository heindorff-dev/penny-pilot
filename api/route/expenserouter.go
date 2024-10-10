package route

import (
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	expense := router.Group("/api/v1/expenses")
	{
		expense.GET("/:id", func(ctx *gin.Context) {
			getExpense(ctx)
		})
		expense.GET("/", func(ctx *gin.Context) {
			getExpenses(ctx)
		})
	}
}

func getExpense(ctx *gin.Context) {
	//var id = ctx.Param("id")
}

func getExpenses(ctx *gin.Context) {
	var name = ctx.Query("name")
	expense, err := service.GetExpense(name)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}
