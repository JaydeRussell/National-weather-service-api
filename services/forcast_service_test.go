package services_test

import (
	"errors"
	"testing"

	"national-weather-service-api/data"
	mock_interfaces "national-weather-service-api/interfaces/mocks"
	"national-weather-service-api/services"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetForcast(t *testing.T) {
	testCases := []struct {
		description string

		lat  float64
		long float64

		mockResponse *data.Forcast
		mockErr      error

		expectedResponse *data.Forcast
		expectedErr      error
	}{
		{
			description: "5 degrees is cold",

			mockResponse: &data.Forcast{
				ShortForcast: "",
				Temperature:  5,
			},
			expectedResponse: &data.Forcast{
				ShortForcast:     "",
				Characterization: "cold",
				Temperature:      5,
			},
		},
		{
			description: "75 degrees is moderate",

			mockResponse: &data.Forcast{
				ShortForcast: "",
				Temperature:  75,
			},
			expectedResponse: &data.Forcast{
				ShortForcast:     "",
				Characterization: "moderate",
				Temperature:      75,
			},
		},
		{
			description: "90 degrees is hot",

			mockResponse: &data.Forcast{
				ShortForcast: "",
				Temperature:  90,
			},
			expectedResponse: &data.Forcast{
				ShortForcast:     "",
				Characterization: "hot",
				Temperature:      90,
			},
		},
		{
			description: "errors on error",

			mockErr:     errors.New("this is an error"),
			expectedErr: errors.New("this is an error"),
		},
	}

	for _, testcase := range testCases {
		// setup
		ctrl := gomock.NewController(t)
		mockGetter := mock_interfaces.NewMockForcastGetter(ctrl)
		forcastService := services.NewForcastService(mockGetter)

		mockGetter.EXPECT().GetForcast(testcase.lat, testcase.long).Times(1).
			Return(testcase.mockResponse, testcase.mockErr)

		// test
		actualResult, actualErr := forcastService.GetForcast(testcase.lat, testcase.long)

		// assert
		if testcase.expectedErr == nil {
			assert.NoError(t, actualErr, testcase.description)
		} else if testcase.expectedErr != nil {
			assert.Error(t, actualErr, testcase.description)
		}

		if testcase.expectedResponse != nil {
			assert.EqualValues(t, *testcase.expectedResponse, *actualResult, testcase.description)
		}
	}
}
