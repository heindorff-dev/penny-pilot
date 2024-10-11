package expense

import (
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetExpense(ctx *gin.Context) {
	id := ctx.Param("id")
	parsed, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	expense, err := service.GetExpenseById(parsed)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

func GetExpenses(ctx *gin.Context) {
	name := ctx.Query("name")
	expense, err := service.GetExpense(name)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}
