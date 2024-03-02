package validation

var(
	ErrIdField = "id"
	ErrNameField = "name"
	ErrDateField = "date"
	ErrPriceField = "price"
	ErrCategoryField = "category"
	ErrMaxAmountField = "maxAmount"
	ErrSpentAmountField = "spentAmount"

	ErrInvalidId = "invalid id"
	ErrInvalidExpenseName = "Expenses need to have a name"
	ErrInvalidExpensePrice = "The price of the expense must be greater than 0"
	ErrInvalidExpenseDate = "Invalid Expense date, it should be in the form of dd/mm/yyyy"
	ErrInvalidExpenseCategory = "invalid Expense category"
	ErrInvalidCategoryName = "The category must have a name"
	ErrCategoryNameExists = "category name already exists"
	ErrInvalidCategoryMaxAmount = "Category max amount must be greater than 0"
	ErrInvalidCategorySpentAmount = "invalid Category spent amount"
	ErrCategoryExists = "category already exists"
	ErrInvalidExpense = "invalid expense"
	ErrInvalidCategory = "invalid category"
)

// var Errors = map[string]string{
// 	ErrInvalidId: "invalid id",
// 	ErrInvalidExpenseName: "Expenses need to have a name",
// 	ErrInvalidExpensePrice: "The price of the expense must be greater than 0",
// 	ErrInvalidExpenseDate: "Invalid Expense date, it should be in the form of dd/mm/yyyy",
// 	ErrInvalidExpenseCategory: "invalid Expense category",
// 	ErrInvalidCategoryName: "The category must have a name",
// 	ErrCategoryNameExists: "category name already exists",
// 	ErrInvalidCategoryMaxAmount: "Category max amount must be greater than 0",
// 	ErrInvalidCategorySpentAmount: "invalid Category spent amount",
// 	ErrCategoryExists: "category already exists",
// }

/*
Como paso los errores en JSON?
Tendria que inclui el success en todas las respuestas para saber
si hubo un error del usuario o no. En caso de que haya fallado el server
si que se devuelve el error en el JSON, pero si es un error del user no hace
falta, directamente se devuelve el status false.
Se me ocurre pasar un array de errores que tengan field y message
*/