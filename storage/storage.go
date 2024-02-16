package storage

import "example/plutonke-server/types"

type Storage interface {
	GetAllExpenses() (*[]types.Expense, error)
	GetExpenseById(uint) (types.Expense, error)
	AddExpense(types.Expense) (types.Expense, error)
	EditExpense(types.Expense) (types.Expense, error)
	DeleteExpense(uint) error

	GetAllCategories() (*[]types.Category, error)
	GetCategoryById(uint) (types.Category, error)
	AddCategory(types.Category) (types.Category, error)
	EditCategory(types.Category) (types.Category, error)	
	DeleteCategory(uint) error

	Close() error
}

