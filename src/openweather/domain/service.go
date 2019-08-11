package domain

import (
	"context"
)

type (
	// Weather state
	Weather struct {
		MainCharacter       string
		Description         string
		IconCode            string
		Temp                int
		Humidity            int
		TempMin             int
		TempMax             int
		WindSpeed           float64
		Cloudiness          int
		LocationName        string
		LocationCountryCode string
	}

	// Service to get the weather
	Service interface {
		GetByCity(ctx context.Context, city string) (Weather, error)
	}
)
