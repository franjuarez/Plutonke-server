package validation

import (
	"example/plutonke-server/storage"
	"example/plutonke-server/types"
	"time"
)

func ValidateExpense(expense types.Expense, storage storage.Storage) bool {
	return validateExpenseName(expense.Name) && validateExpensePrice(expense.Price) &&
		validateExpenseDate(expense.Date) && validateExpenseCategory(expense.CategoryID, storage)
}

func validateExpenseName(name string) bool {
	return len(name) > 0
}

func validateExpensePrice(price float32) bool {
	return price > 0
}

func validateExpenseDate(date int64) bool {
	t := time.Unix(date,  0)
	return !t.IsZero()
}

func validateExpenseCategory(categoryID uint, storage storage.Storage) bool {
	category, _ := storage.GetCategoryById(categoryID)
	result, _ := storage.CheckIfCategoryExists(category)
	return result
}