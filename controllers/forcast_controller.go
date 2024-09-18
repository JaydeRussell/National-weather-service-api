package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"national-weather-service-api/interfaces"
	"national-weather-service-api/tools"

	"github.com/labstack/echo/v4"
)

type ForcastController struct {
	forcaster interfaces.Forcaster
}

func NewForcastController(forcaster interfaces.Forcaster) *ForcastController {
	return &ForcastController{
		forcaster: forcaster,
	}
}

// GetForcast godoc
//
//	@Summary		Fetch a short forcast
//	@Description	Accepts latitude and longitude coordinates then returns a short forcast+ temperature characterization
//	@Tags			forcast
//	@Produce		json
//	@Param			lat		path		float64	true	"latitude"
//	@Param			long	path		float64	true	"longitude"
//	@Success		200		{object}	data.Forcast
//	@Failure		400		{object}	data.HTTPError
//	@Failure		404		{object}	data.HTTPError
//	@Failure		500		{object}	data.HTTPError
//	@Router			/forcast/get/{lat}/{long} [get]
func (w *ForcastController) GetForcast(c echo.Context) error {
	latParam := c.Param("lat")
	longParam := c.Param("long")

	lat, err := strconv.ParseFloat(latParam, 64)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			fmt.Errorf("unable to parse latitude. Must be a single float in decimal degrees: %w", err),
		)
	}

	long, err := strconv.ParseFloat(longParam, 64)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			fmt.Errorf("unable to parse longitude. Must be a single float in decimal degrees: %w", err),
		)
	}

	forcast, forcastErr := w.forcaster.GetForcast(lat, long)
	if forcastErr != nil {
		return tools.HandleError(c, forcastErr)
	}

	return c.JSON(http.StatusOK, forcast)
}
