package utils

import (
	"example/plutonke-server/validation"
	"net/http"

	"github.com/labstack/echo"
)

func SendSuccessResponse[T any](c echo.Context, response T) error {
	return c.JSON(http.StatusOK, response)
}

func SendFailValidationResponse(c echo.Context, response []validation.ValidationError) error {
	return c.JSON(http.StatusBadRequest, response)
}

func SendFailServerResponse(c echo.Context, response map[string]string) error {
	return c.JSON(http.StatusInternalServerError, response)
}

func SendEmptyResponse(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
