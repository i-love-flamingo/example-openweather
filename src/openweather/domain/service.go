package domain

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
)
