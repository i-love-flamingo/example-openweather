package interfaces

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/example-openweather/src/openweather/domain"
)

type (
	// Controller for the openweather routes
	Controller struct {
		responder *web.Responder
	}

	viewData struct {
		City    string
		Weather domain.Weather
	}
)

// Inject dependencies
func (controller *Controller) Inject(responder *web.Responder) {
	controller.responder = responder
}

// Get renders the weather page for the given city
func (controller *Controller) Get(ctx context.Context, r *web.Request) web.Result {
	city := r.Params["city"]
	return controller.responder.Render(
		"weather/weather",
		viewData{
			City: city, Weather: domain.Weather{
				MainCharacter:       "cloudy",
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
			},
		},
	)
}
