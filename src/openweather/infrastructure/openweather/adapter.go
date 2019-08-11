package openweather

import (
	"context"
	"errors"

	"flamingo.me/example-openweather/src/openweather/domain"
)

type (
	// Adapter for openweather
	Adapter struct {
		apiClient *APIClient
	}

	weatherDto struct {
		// use https://mholt.github.io/json-to-go/ to generate it easily
	}
)

var (
	// Check if we really implement the interface during compilation
	_ domain.Service = (*Adapter)(nil)
	// ErrNoWeather is returned if no weather data is available
	ErrNoWeather = errors.New("no weather data")
)

// Inject dependencies
func (a *Adapter) Inject(
	client *APIClient,
) *Adapter {
	a.apiClient = client

	return a
}

// GetByCity returns the weather for the given city
func (a *Adapter) GetByCity(ctx context.Context, city string) (domain.Weather, error) {
	return domain.Weather{}, ErrNoWeather
}

func mapDto(dto *weatherDto) (domain.Weather, error) {
	return domain.Weather{}, nil
}
