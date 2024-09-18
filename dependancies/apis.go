package dependancies

import (
	"national-weather-service-api/apis"
	"national-weather-service-api/constants"
)

type apiDependancies struct {
	nwsAPI *apis.NWSAPI
}

func (d *Dependancies) GetNWSAPI() *apis.NWSAPI {
	if d.apis.nwsAPI == nil {
		d.apis.nwsAPI = apis.NewNWSAPI(
			constants.NWS_APU_URL,
		)
	}

	return d.apis.nwsAPI
}
