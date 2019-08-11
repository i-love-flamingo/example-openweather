package interfaces

import (
	"context"
	"reflect"
	"testing"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/go-test/deep"

	"flamingo.me/example-openweather/src/openweather/application"
	"flamingo.me/example-openweather/src/openweather/domain"
)

type (
	mockService struct{}
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

func TestController_Get(t *testing.T) {
	t.Parallel()
	type fields struct {
		responder *web.Responder
		service   *application.Service
	}
	type args struct {
		ctx context.Context
		r   *web.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   web.Result
	}{
		{
			name: "w/ parameter",
			fields: fields{
				responder: &web.Responder{},
				service:   new(application.Service).Inject(new(mockService), flamingo.NullLogger{}),
			},
			args: args{
				ctx: context.Background(),
				r: &web.Request{
					Params: web.RequestParams{
						"city": "flamingo capital",
					},
				},
			},
			want: &web.RenderResponse{
				DataResponse: web.DataResponse{
					Response: web.Response{
						Status: 200,
						Header: make(map[string][]string),
					},
					Data: viewData{
						City: "flamingo capital",
						Weather: domain.Weather{
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
				},
				Template: "weather/weather",
			},
		},
		{
			name: "w/o parameter",
			fields: fields{
				responder: &web.Responder{},
				service:   new(application.Service).Inject(new(mockService), flamingo.NullLogger{}),
			},
			args: args{
				ctx: context.Background(),
				r:   &web.Request{},
			},
			want: &web.RenderResponse{
				DataResponse: web.DataResponse{
					Response: web.Response{
						Status: 200,
						Header: make(map[string][]string),
					},
					Data: viewData{
						City: "",
						Weather: domain.Weather{
							MainCharacter:       "cloudy",
							Description:         "light intensity drizzle",
							IconCode:            "09d",
							Temp:                280,
							Humidity:            80,
							TempMin:             279,
							TempMax:             281,
							WindSpeed:           4.1,
							Cloudiness:          80,
							LocationName:        "",
							LocationCountryCode: "DE",
						},
					},
				},
				Template: "weather/weather",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := new(Controller).Inject(
				tt.fields.responder,
				tt.fields.service,
			)

			if diff := deep.Equal(controller.Get(tt.args.ctx, tt.args.r), tt.want); diff != nil {
				t.Error("Controller.Get() not as expected: diff: ", diff)
			}
		})
	}
}

func TestController_Data(t *testing.T) {
	type fields struct {
		responder *web.Responder
		service   *application.Service
	}
	type args struct {
		ctx        context.Context
		r          *web.Request
		callParams web.RequestParams
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "w/ parameter",
			fields: fields{
				responder: &web.Responder{},
				service:   new(application.Service).Inject(new(mockService), flamingo.NullLogger{}),
			},
			args: args{
				ctx:        context.Background(),
				r:          nil,
				callParams: web.RequestParams{"city": "flamingo capital"},
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
			name: "w/o parameter",
			fields: fields{
				responder: &web.Responder{},
				service:   new(application.Service).Inject(new(mockService), flamingo.NullLogger{}),
			},
			args: args{
				ctx:        context.Background(),
				r:          nil,
				callParams: nil,
			},
			want: domain.Weather{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &Controller{
				responder: tt.fields.responder,
				service:   tt.fields.service,
			}
			if got := controller.Data(tt.args.ctx, tt.args.r, tt.args.callParams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data() = %v, want %v", got, tt.want)
			}
		})
	}
}
