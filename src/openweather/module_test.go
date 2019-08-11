package openweather_test

import (
	"testing"

	"flamingo.me/dingo"
	"flamingo.me/example-openweather/src/openweather"
)

func TestModule_Configure(t *testing.T) {
	t.Parallel()
	if err := dingo.TryModule(new(openweather.Module)); err != nil {
		t.Error(err)
	}
}
