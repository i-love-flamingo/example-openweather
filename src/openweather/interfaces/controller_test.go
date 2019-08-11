package interfaces

import (
	"context"
	"testing"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/go-test/deep"
)

func TestController_Get(t *testing.T) {
	t.Parallel()
	type fields struct {
		responder *web.Responder
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
					Data: viewData{City: "flamingo capital"},
				},
				Template: "weather/weather",
			},
		},
		{
			name: "w/o parameter",
			fields: fields{
				responder: &web.Responder{},
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
					Data: viewData{City: ""},
				},
				Template: "weather/weather",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &Controller{
				responder: tt.fields.responder,
			}

			if diff := deep.Equal(controller.Get(tt.args.ctx, tt.args.r), tt.want); diff != nil {
				t.Error("Controller.Get() not as expected: diff: ", diff)
			}
		})
	}
}
