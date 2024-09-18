package tools

import (
	"national-weather-service-api/data"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, err error) error {
	if httpError, ok := err.(*data.HTTPError); ok {
		return c.JSON(httpError.StatusCode, httpError)
	}

	return c.JSON(http.StatusInternalServerError, err)
}
