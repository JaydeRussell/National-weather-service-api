package interfaces

import "national-weather-service-api/data"

type ForcastGetter interface {
	GetForcast(lat, long float64) (*data.Forcast, error)
}
