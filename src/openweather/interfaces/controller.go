package interfaces

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/example-openweather/src/openweather/application"
	"flamingo.me/example-openweather/src/openweather/domain"
)

type (
	// Controller for the openweather routes
	Controller struct {
		responder *web.Responder
		service   *application.Service
	}

	viewData struct {
		City    string
		Weather domain.Weather
	}
)

// Inject dependencies
func (controller *Controller) Inject(responder *web.Responder, service *application.Service) *Controller {
	controller.responder = responder
	controller.service = service

	return controller
}

// Get renders the weather page for the given city
func (controller *Controller) Get(ctx context.Context, r *web.Request) web.Result {
	city := r.Params["city"]
	return controller.responder.Render("weather/weather", viewData{City: city, Weather: controller.service.GetWeatherByCityName(ctx, city)})
}
