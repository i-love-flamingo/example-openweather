package application_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"flamingo.me/example-openweather/src/openweather/application"
	"flamingo.me/example-openweather/src/openweather/domain"
	"flamingo.me/flamingo/v3/framework/flamingo"
)

type (
	mockService      struct{}
	mockErrorService struct{}
)

func (*mockService) GetByCity(_ context.Context, city string) (domain.Weather, error) {
	return domain.Weather{
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
	}, nil
}

func (*mockErrorService) GetByCity(ctx context.Context, city string) (domain.Weather, error) {
	return domain.Weather{}, errors.New("a test error")
}

func TestService_GetWeatherByCityName(t *testing.T) {
	type fields struct {
		service domain.Service
		logger  flamingo.Logger
	}
	type args struct {
		ctx  context.Context
		city string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   domain.Weather
	}{
		{
			name: "valid",
			fields: fields{
				service: &mockService{},
				logger:  flamingo.NullLogger{},
			},
			args: args{
				ctx:  nil,
				city: "flamingo capital",
			},
			want: domain.Weather{
				MainCharacter:       "cloudy",
				Description:         "light intensity drizzle",
				IconCode:            "09d",
				Temp:                280,
				Humidity:            80,
				TempMin:             279,
				TempMax:             281,
				WindSpeed:           4.1,
				Cloudiness:          80,
				LocationName:        "flamingo capital",
				LocationCountryCode: "DE",
			},
		},
		{
			name: "error",
			fields: fields{
				service: &mockErrorService{},
				logger:  flamingo.NullLogger{},
			},
			args: args{
				ctx:  nil,
				city: "flamingo capital",
			},
			want: domain.Weather{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(application.Service).Inject(
				tt.fields.service,
				tt.fields.logger,
			)
			if got := s.GetWeatherByCityName(tt.args.ctx, tt.args.city); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetWeatherByCityName() = %v, want %v", got, tt.want)
			}
		})
	}
}
