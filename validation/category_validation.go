package validation

import (
	"example/plutonke-server/storage"
	"example/plutonke-server/types"
)

func ValidateCategory(category types.Category, storage storage.Storage) bool {
	return validateCategoryName(category.Name, storage) &&
		validateCategoryMaxAmount(category.MaxAmount) &&
		validateCategorySpentAmount(category.SpentAmount)
}

func validateCategoryName(name string, storage storage.Storage) bool {
	result, _ := storage.CheckIfCategoryNameExists(name)
	return result
}

func validateCategoryMaxAmount(maxAmount float32) bool {
	return maxAmount > 0
}

func validateCategorySpentAmount(spentAmount float32) bool {
	return true
}