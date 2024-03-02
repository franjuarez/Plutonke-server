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
		*errors = append(*errors, ValidationError{
			Field: ErrNameField,
			Message: ErrInvalidExpenseName,
		})
	}
}

func validateExpensePrice(price float32, errors *[]ValidationError) {
	if price <= 0{
		*errors = append(*errors, ValidationError{
			Field: ErrPriceField,
			Message: ErrInvalidExpensePrice,
		})
	}
}

func validateExpenseDate(date int64, errors *[]ValidationError) {
	t := time.Unix(date,  0)
	if t.IsZero() {
		*errors = append(*errors, ValidationError{
			Field: ErrDateField,
			Message: ErrInvalidExpenseDate,
		})
	}
}

func validateExpenseCategory(categoryID uint, storage storage.Storage, errors *[]ValidationError) {
	category, _ := storage.GetCategoryById(categoryID)
	result, _ := storage.CheckIfCategoryExists(category)
	if !result {
		*errors = append(*errors, ValidationError{
			Field: ErrCategoryField,
			Message: ErrInvalidExpenseCategory,
		})
	}
}