package data

type Forcast struct {
	ShortForcast     string `json:"shortForcast"`
	Characterization string `json:"characterization"`
	Temperature      int    `json:"-"`
}

type NWSForcast struct {
	Properties struct {
		Periods []struct {
			Temperature   int    `json:"temperature"`
			ShortForecast string `json:"shortForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

type NWSPoint struct {
	Properties struct {
		ForecastUrl string `json:"forecast"`
	} `json:"properties"`
}
