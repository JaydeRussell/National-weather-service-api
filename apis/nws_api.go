package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"national-weather-service-api/data"
	"national-weather-service-api/tools"
	"net/http"
)

type NWSAPI struct {
	url    string
	client *http.Client
}

func NewNWSAPI(url string) *NWSAPI {
	return &NWSAPI{
		url:    url,
		client: &http.Client{},
	}
}

func (n *NWSAPI) GetForcast(lat, long float64) (*data.Forcast, error) {
	pointURL, err := n.getPointURL(lat, long)
	if err != nil {
		log.Printf("failed to get point URL %s", err)
		return nil, err
	}

	bodyBytes, statusCode, err := tools.HttpGet(pointURL)
	if err != nil {
		log.Printf("failed to get %s", pointURL)
		return nil, err
	}

	if statusCode != http.StatusOK {
		log.Printf("unexpcted response code %d", statusCode)
		return nil, data.NewHttpError(
			fmt.Sprintf("unexpected response: %s", bodyBytes),
			statusCode,
		)
	}

	body := data.NWSForcast{}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Printf("failed to unmarshal forcast body: %s err: %s", bodyBytes, err)
		return nil, err
	}

	periods := body.Properties.Periods
	if len(periods) == 0 {
		log.Println("no data returned from forcast")
		return nil, errors.New("missing data for coordinates")
	}

	return &data.Forcast{
		ShortForcast: periods[0].ShortForecast,
		Temperature:  periods[0].Temperature,
	}, nil
}

func (n *NWSAPI) getPointURL(lat, long float64) (string, error) {
	url := fmt.Sprintf("%s/points/%f,%f", n.url, lat, long)
	bodyBytes, statusCode, err := tools.HttpGet(url)
	if err != nil {
		log.Printf("failed to get %s. Err: %s", url, err)
		return "", err
	}

	if statusCode != http.StatusOK {
		log.Printf("unexpcted response code %d", statusCode)
		return "", data.NewHttpError(
			fmt.Sprintf("unexpected response: %s", bodyBytes),
			statusCode,
		)
	}

	body := data.NWSPoint{}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Printf("failed to unmarshal forcast body: %s err: %s", bodyBytes, err)
		return "", err
	}

	return body.Properties.ForecastUrl, nil
}
