package dependancies

import "national-weather-service-api/controllers"

type controllerDependancies struct {
	forcastController *controllers.ForcastController
}

func (d *Dependancies) GetForcastController() *controllers.ForcastController {
	if d.controllers.forcastController == nil {
		d.controllers.forcastController = controllers.NewForcastController(
			d.GetForcastService(),
		)
	}

	return d.controllers.forcastController
}
