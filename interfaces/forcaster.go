package interfaces

import "national-weather-service-api/data"

type Forcaster interface {
	GetForcast(lat, long float64) (*data.Forcast, error)
}
