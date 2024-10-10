package service

import (
	"errors"

	"github.com/google/uuid"
)

type ExepenseService struct {
	//expenses []Expense
}

type Expense struct {
	Amount int       `json:"amount"`
	Name   string    `json:"name"`
	Id     uuid.UUID `json:"id"`
}

var expenses = []Expense{
	{
		Amount: 50,
		Name:   "Gas",
		Id:     uuid.New(),
	},
	{
		Amount: 100,
		Name:   "Car",
		Id:     uuid.New(),
	},
}

func GetExpense(name string) (e *Expense, err error) {
	expense, err := getFromDB(name)
	return expense, err
}

func getFromDB(name string) (e *Expense, err error) {
	for _, expense := range expenses {
		if expense.Name == name {
			return &expense, nil
		}
	}
	return nil, errors.New("could not find expense")
}
