package errors

// var(
// 	ErrInvalidId = "invalid id"
// 	ErrInvalidExpenseName = "invalid Expense name"
// 	ErrInvalidExpensePrice = "invalid Expense price"
// 	ErrInvalidExpenseDate = "invalid Expense date"
// 	ErrInvalidExpenseCategory = "invalid Expense category"
// 	ErrInvalidCategoryName = "invalid Category name"
// 	ErrCategoryNameExists = "category name already exists"
// 	ErrInvalidCategoryMaxAmount = "invalid Category max amount"
// 	ErrInvalidCategorySpentAmount = "invalid Category spent amount"
// 	ErrCategoryExists = "category already exists"
// )

var Errors = map[string]string{
	"invalid_id": "invalid id",
	"invalid_expense_name": "Expenses need to have a name",
	"invalid_expense_price": "The price of the expense must be greater than 0",
	"invalid_expense_date": "Invalid Expense date, it should be in the form of dd/mm/yyyy",
	"invalid_expense_category": "invalid Expense category",
	"invalid_category_name": "The category must have a name",
	"category_name_exists": "category name already exists",
	"invalid_category_max_amount": "Category max amount must be greater than 0",
	"invalid_category_spent_amount": "invalid Category spent amount",
	"category_exists": "category already exists",
}

/*
Como paso los errores en JSON?
Tendria que inclui el success en todas las respuestas para saber
si hubo un error del usuario o no. En caso de que haya fallado el server
si que se devuelve el error en el JSON, pero si es un error del user no hace
falta, directamente se devuelve el status false.
Se me ocurre pasar un array de errores que tengan field y message
*/