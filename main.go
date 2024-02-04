package main

import (
	"github.com/labstack/echo"
	"net/http"
	"example/plutonke-server/utils"
	"example/plutonke-server/expense"

)



var expenses = []expense.Expense{
	{Id: "1", Name: "minecraft", Price: 20, Date: "10/03/2020", Category: "Games"},
	{Id: "2", Name: "Cumple de urko 2024", Price: 1500, Date: "10/03/2020", Category: "Gifts"},
	{Id: "3", Name: "Cena con amigos", Price: 50, Date: "10/03/2020", Category: "Food"},
	{Id: "4", Name: "Pool", Price: 1000, Date: "10/03/2024", Category: "Games"},
}

func main() {
	router := echo.New()
	router.GET("/expenses", GetAllExpenses)
	router.POST("/expenses", AddExpense)
	router.GET("/expenses/:id", GetExpense)
	router.PUT("/expenses/:id", EditExpense)
	router.DELETE("/expenses/:id", DeleteExpense)

	router.Start(":8080")
}

func GetAllExpenses(context echo.Context) error {
	return context.JSON(http.StatusOK, expenses)
}


func GetExpense(c echo.Context) error {
	id := c.Param("id")
	expense, err := utils.GetItemById(id, expenses)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, expense)
}

func AddExpense(c echo.Context) error {
	var expense expense.Expense

	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	expenses = append(expenses, expense)

	return c.JSON(http.StatusOK, expense)
}

func EditExpense(c echo.Context) error {
	id := c.Param("id")
	expense, err := utils.GetItemById(id, expenses)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	if err := c.Bind(expense); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, expense)
}



func DeleteExpense(c echo.Context) error {
	id := c.Param("id")

	// Buscar el índice del elemento con el ID dado en la lista
	index, err := utils.GetIndexById(id, expenses)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	// Eliminar el elemento de la lista usando la técnica de "slicing"
	expenses = append(expenses[:index], expenses[index+1:]...)

	return c.JSON(http.StatusOK, map[string]string{"message": "Expense deleted"})
}
	
	
/*
* Obtener todas las expenses ✔️
* Agregar 1 expense ✔️
* Editar 1 expense ✔️
* Borrar 1 expense✔️


* Obtener todas las categorias
* Agregar 1 categoria
* Editar 1 categoria
* Borrar 1 categoria
 */