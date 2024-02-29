package validation

import (
	"example/plutonke-server/storage"
	"example/plutonke-server/types"
	"time"
)

func ValidateExpense(expense types.Expense, storage storage.Storage) []ValidationError{
	errors := []ValidationError{}
	validateExpenseName(expense.Name, &errors)
	validateExpensePrice(expense.Price, &errors)
	validateExpenseDate(expense.Date, &errors)
	validateExpenseCategory(expense.CategoryID, storage, &errors)
	return errors
}

func validateExpenseName(name string, errors *[]ValidationError) {
	if len(name) == 0{
		err := ErrInvalidExpenseName
		*errors = append(*errors, ValidationError{
			Field: err,
			Message: Errors[err],
		})
	}
}

func validateExpensePrice(price float32, errors *[]ValidationError) {
	if price <= 0{
		err := ErrInvalidExpensePrice
		*errors = append(*errors, ValidationError{
			Field: err,
			Message: Errors[err],
		})
	}
}

func validateExpenseDate(date int64, errors *[]ValidationError) {
	t := time.Unix(date,  0)
	if t.IsZero() {
		err := ErrInvalidExpenseDate
		*errors = append(*errors, ValidationError{
			Field: err,
			Message: Errors[err],
		})
	}
}

func validateExpenseCategory(categoryID uint, storage storage.Storage, errors *[]ValidationError) {
	category, _ := storage.GetCategoryById(categoryID)
	result, _ := storage.CheckIfCategoryExists(category)
	if result {
		err := ErrInvalidExpenseCategory
		*errors = append(*errors, ValidationError{
			Field: err,
			Message: Errors[err],
		})
	}
}