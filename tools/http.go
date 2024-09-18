package tools

import (
	"io"
	"national-weather-service-api/constants"
	"net/http"
)

func HttpGet(url string) (body []byte, statusCode int, err error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	request.Header.Add("User-Agent", constants.HTTP_USER_AGENT)

	resp, err := client.Do(request)
	if err != nil {
		return
	}

	statusCode = resp.StatusCode
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
