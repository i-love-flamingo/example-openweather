package infrastructure_test

import (
	"context"
	"reflect"
	"testing"

	"flamingo.me/example-openweather/src/openweather/domain"
	"flamingo.me/example-openweather/src/openweather/infrastructure"
)

func TestFakeservice_GetByCity(t *testing.T) {
	type args struct {
		ctx  context.Context
		city string
	}
	tests := []struct {
		name    string
		f       *infrastructure.Fakeservice
		args    args
		want    domain.Weather
		wantErr bool
	}{
		{
			name: "valid",
			f:    &infrastructure.Fakeservice{},
			args: args{
				ctx:  context.Background(),
				city: "flamingo capital",
			},
			want: domain.Weather{
				MainCharacter:       "fake-cloudy",
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
			wantErr: false,
		},
		{
			name: "error",
			f:    &infrastructure.Fakeservice{},
			args: args{
				ctx:  nil,
				city: "error",
			},
			want:    domain.Weather{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &infrastructure.Fakeservice{}
			got, err := f.GetByCity(tt.args.ctx, tt.args.city)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fakeservice.GetByCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fakeservice.GetByCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
