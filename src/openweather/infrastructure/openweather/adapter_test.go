package openweather

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/testutil"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdapter(t *testing.T) {
	testutil.WithPact(t, "flamingo-helloworld", "openweather", PactTestOpenWeatherAdapter)
}

func PactTestOpenWeatherAdapter(t *testing.T, pact *dsl.Pact) {
	stringFixtureContent := string(`
{
  "coord": {
    "lon": -0.13,
    "lat": 51.51
  },
  "weather": [
    {
      "id": 300,
      "main": "Drizzle",
      "description": "light intensity drizzle",
      "icon": "09d"
    }
  ],
  "base": "stations",
  "main": {
    "temp": 280.32,
    "pressure": 1012,
    "humidity": 81,
    "temp_min": 279.15,
    "temp_max": 281.15
  },
  "visibility": 10000,
  "wind": {
    "speed": 4.1,
    "deg": 80
  },
  "clouds": {
    "all": 90
  },
  "dt": 1485789600,
  "sys": {
    "type": 1,
    "id": 5091,
    "message": 0.0103,
    "country": "GB",
    "sunrise": 1485762037,
    "sunset": 1485794875
  },
  "id": 2643743,
  "name": "London",
  "cod": 200
}
`)

	pact.AddInteraction().
		UponReceiving("An request to openweather").
		Given("Data for London exists").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   "/data/2.5/weather",
			Query:  "appid=APIKEY&q=London&units=metric", // Correct Order!
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Body:   stringFixtureContent,
		})

	if err := pact.Verify(func() error {
		var adapter = new(Adapter).Inject(
			&APIClient{
				baseURL:    fmt.Sprintf("http://%s:%d/data/2.5", pact.Host, pact.Server.Port),
				apiKey:     "APIKEY",
				logger:     flamingo.NullLogger{},
				httpClient: http.DefaultClient,
			},
		)

		weatherInfo, e := adapter.GetByCity(context.Background(), "London")

		require.NoError(t, e)

		assert.Equal(t, "Drizzle", weatherInfo.MainCharacter)
		assert.Equal(t, int(280), weatherInfo.Temp)

		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
