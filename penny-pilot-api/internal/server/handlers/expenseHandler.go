package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type expenseHandler struct {
}

func NewExpenseHandler(r *chi.Mux) {
	handler := &expenseHandler{}

	r.Route("/v1/expenses", func(r chi.Router) {
		r.Get("/", handler.GetExpenses)
	})
}

func (e *expenseHandler) GetExpenses(w http.ResponseWriter, r *http.Request) {

}
