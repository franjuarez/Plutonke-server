package api

import(
	"strconv"
	"example/plutonke-server/types"
	"net/http"
	"github.com/labstack/echo"
	)

// ---------------------------------------Expenses---------------------------------------

func (s *Server) HandleGetAllExpenses(c echo.Context) error {
	expenses, err := s.store.GetAllExpenses()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, expenses)
}

func (s *Server) HandleGetExpenseById(c echo.Context) error {
	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(_id)
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	expense, err := s.store.GetExpenseById(id)
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, expense)
}

func (s *Server) HandleAddExpense(c echo.Context) error {
	var expense types.Expense

	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	expenseAdded, err := s.store.AddExpense(expense)

	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, expenseAdded)
}

func (s *Server) HandleEditExpense(c echo.Context) error{
	var expense types.Expense
	
	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(_id)
	expense.Id = id
	expenseEdited, err := s.store.EditExpense(expense)

	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, expenseEdited)
}

func (s *Server) HandleDeleteExpense(c echo.Context) error{
	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(_id)
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	if err := s.store.DeleteExpense(id); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

// --------------------------------------Categories---------------------------------------

func (s *Server) HandleGetAllCategories(c echo.Context) error{
	categories, err := s.store.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, categories)
}

func (s *Server) HandleGetCategoryById(c echo.Context) error {
	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	
	id := uint(_id)
	category, err := s.store.GetCategoryById(id)
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, category)
}

func (s *Server) HandleAddCategory(c echo.Context) error {
	var category types.Category

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	categoryAdded, err := s.store.AddCategory(category)

	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, categoryAdded)
}

func (s *Server) HandleEditCategory(c echo.Context) error{
	var category types.Category
	
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(_id)
	category.Id = id
	categoryEdited, err := s.store.EditCategory(category)

	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, categoryEdited)
}

func (s *Server) HandleDeleteCategory(c echo.Context) error{
	_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(_id)
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	if err := s.store.DeleteCategory(id); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}