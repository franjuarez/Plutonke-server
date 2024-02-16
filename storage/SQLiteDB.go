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

func (sqldb *SQLiteDatabase) GetExpenseById(id uint) (types.Expense, error) {
	expense := types.Expense{}
	result := sqldb.db.First(&expense, id)
	if result.Error != nil {
		return types.Expense{}, result.Error
	}
	return expense, nil
}

func (sqldb *SQLiteDatabase) AddExpense(expense types.Expense) (types.Expense, error) {
	result := sqldb.db.Create(&expense)
	if result.Error != nil {
		return types.Expense{}, result.Error
	}
	//TODO
	//updatear categories spentAmount
	//category := expense.Category;
	//sqldb.db.Model(&category).Update("spentAmount", category.SpentAmount + expense.Amount)

	return expense, nil
}

func (sqldb *SQLiteDatabase) EditExpense(expense types.Expense) (types.Expense, error) {
	result := sqldb.db.Model(&expense).Updates(expense)
	if result.Error != nil {
		return types.Expense{}, result.Error
	}
	//TODO: update categories spentamount
	return expense, nil
}

func (sqldb *SQLiteDatabase) DeleteExpense(id uint) error {
	result := sqldb.db.Delete(&types.Expense{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sqldb *SQLiteDatabase) GetAllCategories() (*[]types.Category, error) {
	categories := []types.Category{}
	sqldb.db.Find(&categories)
	return &categories, nil
}

func (sqldb *SQLiteDatabase) GetCategoryById(id uint) (types.Category, error) {
	category := types.Category{}
	result := sqldb.db.First(&category, id)
	if result.Error != nil {
		return types.Category{}, result.Error
	}
	return category, nil
}

func (sqldb *SQLiteDatabase) AddCategory(category types.Category) (types.Category, error) {
	result := sqldb.db.Create(&category)
	if result.Error != nil {
		return types.Category{}, result.Error
	}
	//TODO
	//updatear categories spentAmount
	//category := expense.Category;
	//sqldb.db.Model(&category).Update("spentAmount", category.SpentAmount + expense.Amount)
	return category, nil
}

func (sqldb *SQLiteDatabase) EditCategory(category types.Category) (types.Category, error) {
	result := sqldb.db.Model(&category).Updates(category)
	if result.Error != nil {
		return types.Category{}, result.Error
	}
	//TODO: update categories spentamount
	return category, nil
}

func (sqldb *SQLiteDatabase) DeleteCategory(id uint) error {
	result := sqldb.db.Delete(&types.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


func (sqldb *SQLiteDatabase) CheckIfCategoryExists(category types.Category) (bool, error) {
	result := sqldb.db.First(&category, category.Id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (sqldb *SQLiteDatabase) CheckIfCategoryNameExists(name string) (bool, error) {
	result := sqldb.db.Where("name = ?", name)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (sqldb *SQLiteDatabase) Close() error {
	if err := sqldb.db.Close(); err != nil {
		return err
	}
	return nil
}
