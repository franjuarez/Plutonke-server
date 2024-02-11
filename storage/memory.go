package storage

import (
	"errors"
	"example/plutonke-server/types"
	"example/plutonke-server/utils"
)

// ------------------------------------Stuct---------------------------------------------

type MemoryStorage struct {
	expenses   []types.Expense
	categories []types.Category
}

// ----------------------------Inicializacion con Datos----------------------------------

var testExpenses = []types.Expense{
	{Id: "1", Name: "Minecraft", Price: 2200, Date: "10/03/2020", Category: "Games"},
	{Id: "2", Name: "Cumple de urko 2024", Price: 1500, Date: "10/03/2020", Category: "Gifts"},
	{Id: "3", Name: "Cena con amigos", Price: 11000, Date: "10/03/2020", Category: "Food"},
	{Id: "4", Name: "Pool", Price: 1000, Date: "10/03/2024", Category: "Games"},
}

var testCategories = []types.Category{
	{Id: "1", Name: "Games", MaxAmount: 3000.5, SpentAmount: 3200.4},
	{Id: "2", Name: "Gifts", MaxAmount: 7000.823, SpentAmount: 1500},
	{Id: "3", Name: "Food", MaxAmount: 50000.324, SpentAmount: 1100},
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		expenses:   testExpenses,
		categories: testCategories,
	}
}

// ---------------------------------Expenses---------------------------------------------

func (ms *MemoryStorage) GetAllExpenses() (*[]types.Expense, error) {
	return &ms.expenses, nil
}

func (ms *MemoryStorage) GetExpenseById(id string) (types.Expense, error) {
	expense, err := utils.GetItemById(id, ms.expenses)
	if err != nil {
		return types.Expense{}, err
	}
	return *expense, nil
}

func (ms *MemoryStorage) AddExpense(expense types.Expense) (types.Expense, error) {
	if isValid := types.ValidateExpense(expense, &ms.categories); !isValid {
		return types.Expense{}, errors.New("invalid expense")
	}
	ms.expenses = append(ms.expenses, expense)
	types.UpdateCategorySpentAmount(expense.Category, &ms.categories, expense.Price)
	return expense, nil
}

func (ms *MemoryStorage) EditExpense(editedExpense types.Expense) (types.Expense, error) {
	if isValid := types.ValidateExpense(editedExpense, &ms.categories); !isValid {
		return types.Expense{}, errors.New("invalid expense")
	}

	id := editedExpense.Id
	index, err := utils.GetIndexById(id, ms.expenses)

	if err != nil {
		return types.Expense{}, err
	}

	difference := editedExpense.Price - ms.expenses[index].Price
	ms.expenses[index] = editedExpense
	types.UpdateCategorySpentAmount(editedExpense.Category, &ms.categories, difference)
	return editedExpense, nil
}

func (ms *MemoryStorage) DeleteExpense(id string) error {
	index, err := utils.GetIndexById(id, ms.expenses)
	if err != nil {
		return err
	}

	ms.expenses = append(ms.expenses[:index], ms.expenses[index+1:]...)
	difference :=  ms.expenses[index].Price * -1
	types.UpdateCategorySpentAmount(ms.expenses[index].Category, &ms.categories, difference)
	return nil
}

//----------------------------Categories-------------------------------------------------

func (ms *MemoryStorage) GetAllCategories() (*[]types.Category, error) {
	return &ms.categories, nil
}

func (ms *MemoryStorage) GetCategoryById(id string) (types.Category, error) {
	category, err := utils.GetItemById(id, ms.categories)
	if err != nil {
		return types.Category{}, err
	}
	return *category, nil
}

func (ms *MemoryStorage) AddCategory(category types.Category) (types.Category, error) {
	if isValid := types.ValidateCategory(category); !isValid {
		return types.Category{}, errors.New("invalid category")
	}
	ms.categories = append(ms.categories, category)
	return category, nil
}

func (ms *MemoryStorage) EditCategory(category types.Category) (types.Category, error) {
	if isValid := types.ValidateCategory(category); !isValid {
		return types.Category{}, errors.New("invalid category")
	}

	id := category.Id
	index, err := utils.GetIndexById(id, ms.categories)

	if err != nil {
		return types.Category{}, err
	}

	ms.categories[index] = category
	return category, nil
}

func (ms *MemoryStorage) DeleteCategory(id string) error {
	index, err := utils.GetIndexById(id, ms.categories)
	if err != nil {
		return err
	}

	ms.categories = append(ms.categories[:index], ms.categories[index+1:]...)
	return nil
}
