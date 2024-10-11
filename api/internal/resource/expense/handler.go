package expense

import (
	"api/internal/database"
	"api/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Expense struct {
	Amount int       `json:"amount"`
	Name   string    `json:"name"`
	Id     uuid.UUID `json:"id"`
}

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

func GetExpenses(c *gin.Context) {
	var (
		queryParamName = c.Query("name")
		expenses       = []Expense{}
		id             uuid.UUID
		name           string
		amount         int
		err            error
	)

	query := "SELECT id, name, amount FROM Expense where name LIKE %" + queryParamName + "%"

	fmt.Println("query: ", query)

	scanner := database.Session.Query(query).WithContext(c).Iter().Scanner()

	for scanner.Next() {
		err = scanner.Scan(&id, &name, &amount)
		if err != nil {
			log.Fatal(err)
		}
		expense := Expense{
			Id:     id,
			Name:   name,
			Amount: amount,
		}
		expenses = append(expenses, expense)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}
