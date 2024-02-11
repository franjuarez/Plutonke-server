package types

import "time"

const DATE_FORMAT = "02/01/2006"

type Expense struct {
	Id       	string 		`json:"id"`
	Name     	string 		`json:"name"`
	Price    	float32		`json:"price"`
	Date     	string 		`json:"date"`
	Category 	string 		`json:"category"`
}

func (e Expense) GetId() string {
	return e.Id
}

func ValidateExpense(expense *Expense, categories []*Category) bool {
	return validateName(expense.Name) && validatePrice(expense.Price) &&
		validateDate(expense.Date) && validateCategory(expense.Category, categories)
}

func validateName(name string) bool {
	return len(name) > 0
}

func validatePrice(price float32) bool {
	return true
}

func validateDate(date string) bool {
	_, err := time.Parse(DATE_FORMAT, date)
	return err == nil
}

func validateCategory(categoryName string, categories []*Category) bool {
	for _, category := range categories {
		if category.Name == categoryName {
			return true
		}
	}
	return false
}