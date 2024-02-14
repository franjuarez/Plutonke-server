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
	db, err := gorm.Open("sqlite3", "plutonke_database.db")
	if err != nil {
		db.Close()
		panic(err)
	}

	// AutoMigrate will create the table if it does not exist
	db.AutoMigrate(&types.Category{})
	db.AutoMigrate(&types.Expense{})

	return &SQLiteDatabase{
		db: db,
	}
}

func (sqldb *SQLiteDatabase) GetAllExpenses() (*[]types.Expense, error) {
	expenses := []types.Expense{}
	sqldb.db.Find(&expenses)
	return &expenses, nil
}

func (sqldb *SQLiteDatabase) GetExpenseById(string) (types.Expense, error) {
	return types.Expense{}, nil
}
func (sqldb *SQLiteDatabase) AddExpense(types.Expense) (types.Expense, error) {
	return types.Expense{}, nil
}
func (sqldb *SQLiteDatabase) EditExpense(types.Expense) (types.Expense, error) {
	return types.Expense{}, nil
 }
func (sqldb *SQLiteDatabase) DeleteExpense(string) error {
	return nil

}
func (sqldb *SQLiteDatabase) GetAllCategories() (*[]types.Category, error) {
	categories := []types.Category{}
	sqldb.db.Find(&categories)
	return &categories, nil
}

func (sqldb *SQLiteDatabase) GetCategoryById(string) (types.Category, error) {
	return types.Category{}, nil
}
func (sqldb *SQLiteDatabase) AddCategory(types.Category) (types.Category, error) {
	return types.Category{}, nil

}
func (sqldb *SQLiteDatabase) EditCategory(types.Category) (types.Category, error)	{
	return types.Category{}, nil
}
func (sqldb *SQLiteDatabase) DeleteCategory(string) error {
	return nil
}

// GetExpenseById(string) (types.Expense, error)
// AddExpense(types.Expense) (types.Expense, error)
// EditExpense(types.Expense) (types.Expense, error)
// DeleteExpense(string) error

func (sqldb *SQLiteDatabase) Close() error {
	if err := sqldb.db.Close(); err != nil{
		return err
	}
	return nil
}