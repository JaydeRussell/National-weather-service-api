package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"national-weather-service-api/controllers"
	"national-weather-service-api/data"
	mock_interfaces "national-weather-service-api/interfaces/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetForcast(t *testing.T) {
	testCases := []struct {
		description string

		latParam  string
		longParam string
		lat       float64
		long      float64

		mockCallsExpected int
		mockResponse      *data.Forcast
		mockErr           error

		expectedErr      error
		expectedStatus   int
		expectedResponse string
	}{
		{
			description: "invalid latitude will fail with a 400",
			latParam:    "thiswillfail",
			longParam:   "8.0",

			expectedStatus: 400,
		},
		{
			description: "invalid longitude will fail with a 400",
			latParam:    "8",
			longParam:   "thiswillfail",

			expectedStatus: 400,
		},
		{
			description: "invalid longitude will fail with a 400",
			latParam:    "8",
			longParam:   "8",

			lat:  8,
			long: 8,

			mockCallsExpected: 1,
			mockResponse: &data.Forcast{
				ShortForcast:     "short forcast",
				Characterization: "moderate",
				Temperature:      75,
			},

			expectedResponse: "{\"shortForcast\":\"short forcast\",\"characterization\":\"moderate\"}\n",
			expectedStatus:   200,
		},
		{
			description: "api returning a 404 returns a 404",
			latParam:    "8",
			longParam:   "8",

			lat:  8,
			long: 8,

			mockCallsExpected: 1,
			mockErr:           &data.HTTPError{StatusCode: 404, Message: "not found"},

			expectedStatus: 404,
		},
	}

	for _, testCase := range testCases {
		// setup
		ctrl := gomock.NewController(t)
		mockForcaster := mock_interfaces.NewMockForcaster(ctrl)
		forcastController := controllers.NewForcastController(mockForcaster)

		mockForcaster.EXPECT().GetForcast(testCase.lat, testCase.long).Times(testCase.mockCallsExpected).
			Return(testCase.mockResponse, testCase.mockErr)

		e := echo.New()
		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/forcast/get/%s/%s", testCase.latParam, testCase.longParam),
			nil,
		)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetParamNames("lat", "long")
		c.SetParamValues(testCase.latParam, testCase.longParam)

		// test
		actualErr := forcastController.GetForcast(c)

		// assert
		if testCase.expectedErr == nil {
			assert.NoError(t, actualErr, testCase.description)
		} else if testCase.expectedErr != nil {
			assert.Error(t, actualErr, testCase.description)
		}

		if testCase.expectedStatus != 0 {
			assert.Equal(t, testCase.expectedStatus, rec.Result().StatusCode)
		}

		if testCase.expectedResponse != "" {
			assert.EqualValues(t, testCase.expectedResponse, rec.Body.String(), testCase.description)
		}
	}
}
