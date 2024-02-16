package validation

import (
	"example/plutonke-server/storage"
	"example/plutonke-server/types"
	"time"
)

const DATE_FORMAT = "02/01/2006"


func ValidateExpense(expense types.Expense, storage storage.Storage) bool {
	return validateExpenseName(expense.Name) && validateExpensePrice(expense.Price) &&
		validateExpenseDate(expense.Date) && validateExpenseCategory(expense.Category, storage)
}

func validateExpenseName(name string) bool {
	return len(name) > 0
}

func validateExpensePrice(price float32) bool {
	return price > 0
}

func validateExpenseDate(date string) bool {
	_, err := time.Parse(DATE_FORMAT, date)
	return err == nil
}

func validateExpenseCategory(category types.Category, storage storage.Storage) bool {
	result, _ := storage.CheckIfCategoryExists(category)
	return result
}