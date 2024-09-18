package services

import (
	"national-weather-service-api/data"
	"national-weather-service-api/interfaces"
)

type ForcastService struct {
	forcastGetter interfaces.ForcastGetter
}

func NewForcastService(forcastGetter interfaces.ForcastGetter) *ForcastService {
	return &ForcastService{
		forcastGetter: forcastGetter,
	}
}

func (f *ForcastService) GetForcast(lat, long float64) (*data.Forcast, error) {
	forcast, err := f.forcastGetter.GetForcast(lat, long)
	if err != nil {
		return nil, err
	}

	forcast.Characterization = determineCharacterization(forcast.Temperature)

	return forcast, nil
}

func determineCharacterization(temp int) string {
	switch {
	case temp < 60:
		return "cold"
	case temp < 80:
		return "moderate"
	default:
		return "hot"
	}
}
