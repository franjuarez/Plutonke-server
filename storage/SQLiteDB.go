package storage

import (
	"example/plutonke-server/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLiteDatabase struct {
	db *gorm.DB
}

func NewSQLiteStorage() *SQLiteDatabase {
	db, err := gorm.Open("sqlite3", "plutonke_database.db")
	if err != nil {
		db.Close()
		panic(err)
	}

	db.AutoMigrate(&types.Category{})
	db.AutoMigrate(&types.Expense{})

	return &SQLiteDatabase{
		db: db,
	}
}

func (sqldb *SQLiteDatabase) GetAllExpenses() (*[]types.Expense, error) {
	expenses := []types.Expense{}
	if result := sqldb.db.Find(&expenses); result.Error != nil {
		return nil, result.Error
	}
	return &expenses, nil
}

func (sqldb *SQLiteDatabase) GetExpenseById(id uint) (types.Expense, error) {
	expense := types.Expense{}
	if result := sqldb.db.First(&expense, id); result.Error != nil {
		return types.Expense{}, result.Error
	}
	return expense, nil
}

func (sqldb *SQLiteDatabase) AddExpense(expense types.Expense) (types.Expense, error) {
	if result := sqldb.db.Create(&expense); result.Error != nil {
		return types.Expense{}, result.Error
	}
	categoryID := expense.CategoryID
	sqldb.db.Model(&types.Category{}).Where("id = ?", categoryID).
		Update("spent_amount", gorm.Expr("spent_amount + ?", expense.Price))

	return expense, nil
}

func (sqldb *SQLiteDatabase) EditExpense(editedExpense types.Expense) (types.Expense, error) {
	var oldCategory types.Category
	var newCategory types.Category
	var oldExpense types.Expense

	if result := sqldb.db.First(&oldExpense, editedExpense.Id); result.Error != nil {
		return types.Expense{}, result.Error
	}
	if result := sqldb.db.First(&oldCategory, oldExpense.CategoryID); result.Error != nil {
		return types.Expense{}, result.Error
	}
	if result := sqldb.db.Model(&editedExpense).Updates(editedExpense); result.Error != nil {
		return types.Expense{}, result.Error
	}
	if result := sqldb.db.First(&newCategory, editedExpense.CategoryID); result.Error != nil {
		return types.Expense{}, result.Error
	}

	if oldCategory.Id == newCategory.Id {
		difference := editedExpense.Price - oldExpense.Price
		UpdateCategorySpentAmount(oldCategory, sqldb, difference)
	} else {
		UpdateCategorySpentAmount(oldCategory, sqldb, oldExpense.Price*-1)
		UpdateCategorySpentAmount(newCategory, sqldb, editedExpense.Price)
	}
	return editedExpense, nil
}

func (sqldb *SQLiteDatabase) DeleteExpense(id uint) error {
	var expense types.Expense
	if result := sqldb.db.First(&expense, id); result.Error != nil {
		return result.Error
	}

	if result := sqldb.db.Model(&types.Category{}).Where("id = ?", expense.CategoryID).
		Update("spent_amount", gorm.Expr("spent_amount - ?", expense.Price)); result.Error != nil {
		return result.Error
	}

	if result := sqldb.db.Delete(&types.Expense{}, id); result.Error != nil {
		return result.Error
	}

	return nil
}

func (sqldb *SQLiteDatabase) GetAllCategories() (*[]types.Category, error) {
	var categories []types.Category
	if result := sqldb.db.Find(&categories); result.Error != nil {
		return nil, result.Error
	}

	return &categories, nil
}

func (sqldb *SQLiteDatabase) GetCategoryById(id uint) (types.Category, error) {
	var category types.Category
	if result := sqldb.db.First(&category, id); result.Error != nil {
		return types.Category{}, result.Error
	}

	return category, nil
}

func (sqldb *SQLiteDatabase) AddCategory(category types.Category) (types.Category, error) {
	if result := sqldb.db.Create(&category); result.Error != nil {
		return types.Category{}, result.Error
	}
	return category, nil
}

func (sqldb *SQLiteDatabase) EditCategory(category types.Category) (types.Category, error) {
	if result := sqldb.db.Model(&category).Updates(category); result.Error != nil {
		return types.Category{}, result.Error
	}
	return category, nil
}

func (sqldb *SQLiteDatabase) DeleteCategory(id uint) error {
	if result := sqldb.db.Delete(&types.Category{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (sqldb *SQLiteDatabase) CheckIfCategoryExists(category types.Category) (bool, error) {
	if result := sqldb.db.First(&category, category.Id); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (sqldb *SQLiteDatabase) CheckIfCategoryNameExists(name string) (bool, error) {
	if result := sqldb.db.Where("name = ?", name); result.Error != nil {
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
