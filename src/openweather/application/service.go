package application

import (
	"context"

	"flamingo.me/flamingo/v3/framework/flamingo"

	"flamingo.me/example-openweather/src/openweather/domain"
)

// Service to get the weather from
type Service struct {
	service domain.Service
	logger  flamingo.Logger
}

// Inject dependencies
func (s *Service) Inject(service domain.Service, logger flamingo.Logger) *Service {
	s.service = service
	s.logger = logger

	return s
}

// GetWeatherByCityName returns the weather for the given city
func (s *Service) GetWeatherByCityName(ctx context.Context, city string) domain.Weather {
	weather, err := s.service.GetByCity(ctx, city)
	if err != nil {
		s.logger.Error(err)
	}
	return weather
}
