package infrastructure

import (
	"context"
	"errors"

	"flamingo.me/example-openweather/src/openweather/domain"
)

// Fakeservice for weather data
type Fakeservice struct{}

// Check if we really implement the interface during compilation
var _ domain.Service = (*Fakeservice)(nil)

// GetByCity returns a fixed faked weather object
func (f *Fakeservice) GetByCity(ctx context.Context, city string) (domain.Weather, error) {
	if city == "error" {
		return domain.Weather{}, errors.New("error while fetching weather")
	}

	return domain.Weather{
		MainCharacter:       "fake-cloudy",
		Description:         "light intensity drizzle",
		IconCode:            "09d",
		Temp:                280,
		Humidity:            80,
		TempMin:             279,
		TempMax:             281,
		WindSpeed:           4.1,
		Cloudiness:          80,
		LocationName:        city,
		LocationCountryCode: "DE",
	}, nil
}
