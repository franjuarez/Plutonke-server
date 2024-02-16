package types

type Category struct {
	Id       		uint		`json:"id" gorm:"primaryKey;autoIncrement"`
	Name     		string 		`json:"name"        gorm:"not null"`
	MaxAmount		float32		`json:"maxAmount"   gorm:"not null"`
	SpentAmount 	float32		`json:"spentAmount" gorm:"not null"`
}

func (c Category) GetId() uint{
	return c.Id
}

func UpdateCategorySpentAmount(categoryName string, categories *[]Category, difference float32) {
	var category *Category
	for i, currentCategory := range *categories {
		if currentCategory.Name == categoryName {
			category = &(*categories)[i]
			break
		}
	}
	category.SpentAmount += difference
}

func ValidateCategory(category Category, categories *[]Category) bool {
	return validateCategoryName(category.Name, categories) && 
		validateCategoryMaxAmount(category.MaxAmount) &&
		validateCategorySpentAmount(category.SpentAmount)
}

func validateCategoryName(name string, categories *[]Category) bool{
	for _, category := range *categories{
		if category.Name == name{
			return false
		}
	}
	return len(name) > 0
}

func validateCategoryMaxAmount(maxAmount float32) bool {
	return maxAmount > 0
}

func validateCategorySpentAmount(spentAmount float32) bool {
	return true
}
