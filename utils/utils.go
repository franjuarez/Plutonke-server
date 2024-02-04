package utils

import(
	"errors"
	"example/plutonke-server/expense"
)

// Función auxiliar para obtener el índice del elemento con el ID dado
func GetIndexById(id string, expenses []expense.Expense) (int, error) {
	for i, expense := range expenses {
		if id == expense.Id {
			return i, nil
		}
	}

	return -1, errors.New("expense not found")
}

func GetItemById(id string, expenses []expense.Expense) (*expense.Expense, error) {
	for i, item := range expenses {
		if item.Id == id {
			return &expenses[i], nil
		}
	}
	return nil, errors.New("no se encontro la Expense")
}

