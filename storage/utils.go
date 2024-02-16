package storage

import "example/plutonke-server/types"

func UpdateCategorySpentAmount(category types.Category, storage Storage, difference float32) {
	category.SpentAmount += difference
	storage.EditCategory(category)
}