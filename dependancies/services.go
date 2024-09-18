package dependancies

import "national-weather-service-api/services"

type serviceDependancies struct {
	forcastService *services.ForcastService
}

func (d *Dependancies) GetForcastService() *services.ForcastService {
	if d.services.forcastService == nil {
		d.services.forcastService = services.NewForcastService(
			d.GetNWSAPI(),
		)
	}

	return d.services.forcastService
}
