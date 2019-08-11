package openweather

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/example-openweather/src/openweather/domain"
	"flamingo.me/example-openweather/src/openweather/infrastructure"
	"flamingo.me/example-openweather/src/openweather/interfaces"
)

// Module Basic struct
type Module struct{}

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))

	injector.Bind(new(domain.Service)).To(infrastructure.Fakeservice{})
}

type routes struct {
	controller *interfaces.Controller
}

// Inject dependencies
func (r *routes) Inject(controller *interfaces.Controller) {
	r.controller = controller
}

// Routes definition for the module
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.HandleGet("openweather.detail", r.controller.Get)
	registry.Route("/weather/:city", "openweather.detail")
}
