package storage

//TODO: ACORDARSE DE PRIMERO MIGRAR IDS A INT O VER COMO SE TRATAN EN LA BDD

import (
	"example/plutonke-server/types"
	//"example/plutonke-server/utils"
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLiteDatabase struct {
	db *gorm.DB
}



func NewSQLiteStorage() *SQLiteDatabase {
	// Open the SQLite database connection
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// AutoMigrate will create the table if it does not exist
	db.AutoMigrate(&types.Category{})
	db.AutoMigrate(&types.Expense{})

	return &SQLiteDatabase{
		db: db,
	}
}

func (db *SQLiteDatabase) GetAllExpenses() (*[]types.Expense, error) {
	expenses := []types.Expense{}
	db.db.Find(&expenses)
	return &expenses, nil
}

func (db *SQLiteDatabase) GetExpenseById(string) (types.Expense, error) {
	return types.Expense{}, nil
}
func (db *SQLiteDatabase) AddExpense(types.Expense) (types.Expense, error) {
	return types.Expense{}, nil
}
func (db *SQLiteDatabase) EditExpense(types.Expense) (types.Expense, error) {
	return types.Expense{}, nil
 }
func (db *SQLiteDatabase) DeleteExpense(string) error {
	return nil

}
func (db *SQLiteDatabase) GetAllCategories() (*[]types.Category, error) {
	categories := []types.Category{}
	db.db.Find(&categories)
	return &categories, nil
}

func (db *SQLiteDatabase) GetCategoryById(string) (types.Category, error) {
	return types.Category{}, nil
}
func (db *SQLiteDatabase) AddCategory(types.Category) (types.Category, error) {
	return types.Category{}, nil

}
func (db *SQLiteDatabase) EditCategory(types.Category) (types.Category, error)	{
	return types.Category{}, nil
}
func (db *SQLiteDatabase) DeleteCategory(string) error {
	return nil
}

// GetExpenseById(string) (types.Expense, error)
// AddExpense(types.Expense) (types.Expense, error)
// EditExpense(types.Expense) (types.Expense, error)
// DeleteExpense(string) error
