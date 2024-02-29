package api

import (
	"example/plutonke-server/types"
	"example/plutonke-server/utils"
	"example/plutonke-server/validation"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ---------------------------------------Expenses---------------------------------------

func (s *Server) HandleGetAllExpenses(c echo.Context) error {
	expenses, err := s.store.GetAllExpenses()
	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendSuccessResponse(c, expenses)
}


func (s *Server) HandleGetExpenseById(c echo.Context) error {
	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(_id)
	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(), //invalid id
		})
	}

	expense, err := s.store.GetExpenseById(id)
	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendSuccessResponse(c, expense)
}

func (s *Server) HandleAddExpense(c echo.Context) error {
	var expense types.Expense

	if err := c.Bind(&expense); err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	
	if result := validation.ValidateExpense(expense, s.store); !result{
		return utils.SendFailResponse(c, map[string]string{
			"error": "invalid expense",
		})
	}

	expenseAdded, err := s.store.AddExpense(expense)

	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendSuccessResponse(c, expenseAdded)
}

func (s *Server) HandleEditExpense(c echo.Context) error{
	var expense types.Expense
	
	if err := c.Bind(&expense); err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}

	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	id := uint(_id)
	expense.Id = id

	if result := validation.ValidateExpense(expense, s.store); !result{
		return utils.SendFailResponse(c, map[string]string{
			"error": "invalid expense",
		})
	}

	expenseEdited, err := s.store.EditExpense(expense)

	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendSuccessResponse(c, expenseEdited)
}

func (s *Server) HandleDeleteExpense(c echo.Context) error{
	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	
	id := uint(_id)
	if err := s.store.DeleteExpense(id); err != nil{
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendEmptyResponse(c)
}

// --------------------------------------Categories---------------------------------------

func (s *Server) HandleGetAllCategories(c echo.Context) error{
	categories, err := s.store.GetAllCategories()
	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendSuccessResponse(c, categories)
}

func (s *Server) HandleGetCategoryById(c echo.Context) error {
	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	
	id := uint(_id)
	category, err := s.store.GetCategoryById(id)
	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendSuccessResponse(c, category)
}

func (s *Server) HandleAddCategory(c echo.Context) error {
	var category types.Category

	if err := c.Bind(&category); err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}

	if result := validation.ValidateCategory(category, s.store); !result{
		return utils.SendFailResponse(c, map[string]string{
			"error": "Invalid Category",
		})
	}

	categoryAdded, err := s.store.AddCategory(category)

	if err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendSuccessResponse(c, categoryAdded)
}

func (s *Server) HandleEditCategory(c echo.Context) error{
	var category types.Category
	
	if err := c.Bind(&category); err != nil {
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}

	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil{
		return utils.SendFailResponse(c, map[string]string{
			"error": "Invalid id",
		})
	}
	id := uint(_id)
	category.Id = id

	if result := validation.ValidateCategory(category, s.store); !result{
		return utils.SendFailResponse(c, map[string]string{
			"error": "Invalid Category",
		})
	}

	categoryEdited, err := s.store.EditCategory(category)

	if err != nil{
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return utils.SendSuccessResponse(c, categoryEdited)
}

func (s *Server) HandleDeleteCategory(c echo.Context) error{
	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(_id)
	if err != nil{
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}

	if err := s.store.DeleteCategory(id); err != nil{
		return utils.SendFailResponse(c, map[string]string{
			"error": err.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}