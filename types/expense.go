package types

import (
	"time"
)

const DATE_FORMAT = "02/01/2006"

type Expense struct {
	Id       	string 		`json:"id"       gorm:"primary_key"`
	Name     	string 		`json:"name"     gorm:"not null"`
	Price    	float32		`json:"price"    gorm:"not null"`
	Date     	string 		`json:"date"     gorm:"not null"`
	Category 	string 		`json:"category" gorm:"not null"`
}

func (e Expense) GetId() string {
	return e.Id
}

func ValidateExpense(expense Expense, categories *[]Category) bool {
	return validateExpenseName(expense.Name) && validateExpensePrice(expense.Price) &&
		validateExpenseDate(expense.Date) && validateExpenseCategory(expense.Category, categories)
}

func validateExpenseName(name string) bool {
	return len(name) > 0
}

func validateExpensePrice(price float32) bool {
	return true
}

func validateExpenseDate(date string) bool {
	_, err := time.Parse(DATE_FORMAT, date)
	return err == nil
}

func validateExpenseCategory(categoryName string, categories *[] Category) bool {
	for _, category := range *categories {
		if category.Name == categoryName {
			return true
		}
	}
	return false
}