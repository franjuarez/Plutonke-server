package storage

import "example/plutonke-server/types"

type Storage interface {
	GetAllExpenses() (*[]types.Expense, error)
	GetExpenseById(string) (types.Expense, error)
	AddExpense(types.Expense) (types.Expense, error)
	EditExpense(types.Expense) (types.Expense, error)
	DeleteExpense(string) error

	GetAllCategories() (*[]types.Category, error)
	GetCategoryById(string) (types.Category, error)
	AddCategory(types.Category) (types.Category, error)
	EditCategory(types.Category) (types.Category, error)	
	DeleteCategory(string) error

	Close() error
}

