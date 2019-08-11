package interfaces

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	// Controller for the openweather routes
	Controller struct {
		responder *web.Responder
	}

	viewData struct {
		City string
	}
)

// Inject dependencies
func (controller *Controller) Inject(responder *web.Responder) {
	controller.responder = responder
}

// Get renders the weather page for the given city
func (controller *Controller) Get(ctx context.Context, r *web.Request) web.Result {
	city := r.Params["city"]
	return controller.responder.Render("weather/weather", viewData{City: city})
}
