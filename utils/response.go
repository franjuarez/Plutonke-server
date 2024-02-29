package utils

import(
	"github.com/labstack/echo"
	"net/http"
)

type Response[T any] struct {
	Success bool `json:"success"`
	Body    T    `json:"body"`
}

func SendSuccessResponse[T any](c echo.Context, response T) error{
	return c.JSON(http.StatusOK, Response[T]{
		Success: true,
		Body: response,
	})
}

func SendFailResponse[T any](c echo.Context, response T) error{
	return c.JSON(http.StatusOK, Response[T]{
		Success: false,
		Body: response,
	})
}

func SendEmptyResponse(c echo.Context) error{
	return c.NoContent(http.StatusNoContent)
}