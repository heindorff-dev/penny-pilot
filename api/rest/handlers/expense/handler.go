package expense

import (
	"api/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type Expense struct {
	Amount int        `json:"amount"`
	Name   string     `json:"name"`
	Id     gocql.UUID `json:"id"`
}

func GetExpense(c *gin.Context) {
	var (
		id     gocql.UUID
		name   string
		amount int
	)

	paramId := c.Param("id")
	parsedId, err := gocql.ParseUUID(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	query := "SELECT id, name, amount FROM Expense where id = ? LIMIT 1"
	err = database.Session.Query(query, parsedId).WithContext(c).Consistency(gocql.One).Scan(&id, &name, &amount)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	expense := Expense{
		Id:     id,
		Name:   name,
		Amount: amount,
	}

	c.JSON(http.StatusOK, expense)
}

func GetExpenses(c *gin.Context) {
	var (
		queryParamName = c.Query("name")
		expenses       = []Expense{}
		id             gocql.UUID
		name           string
		amount         int
		err            error
	)

	query := "SELECT id, name, amount FROM Expense where name LIKE '%?%'"

	scanner := database.Session.Query(query, queryParamName).WithContext(c).Iter().Scanner()

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
