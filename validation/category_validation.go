package validation

import (
	"example/plutonke-server/storage"
	"example/plutonke-server/types"
)

func ValidateCategory(category types.Category, storage storage.Storage) []ValidationError  {
	errors := []ValidationError{}
	validateCategoryName(category.Name, storage, &errors)
	validateCategoryMaxAmount(category.MaxAmount, &errors)
	// validateCategorySpentAmount(category.SpentAmount &errors)
	return errors
}

func validateCategoryName(name string, storage storage.Storage, errors *[]ValidationError) {
	if len(name) == 0 {
		*errors = append(*errors, ValidationError{
			Field: ErrNameField,
			Message: ErrInvalidCategoryName,
		})
	}

	result, _ := storage.CheckIfCategoryNameExists(name)
	
	if result {
		*errors = append(*errors, ValidationError{
			Field: ErrNameField,
			Message: ErrCategoryNameExists,
		})
	}
}

func validateCategoryMaxAmount(maxAmount float32, errors *[]ValidationError) {
	if maxAmount <= 0 {
		*errors = append(*errors, ValidationError{
			Field: ErrMaxAmountField,
			Message: ErrInvalidCategoryMaxAmount,
		})
	
	}
}

// func validateCategorySpentAmount(spentAmount float32, errors *[]ValidationError) {
// 	return
// }