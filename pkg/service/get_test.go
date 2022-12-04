package service

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestService_Get(t *testing.T) {
	type fields struct {
		Logger echo.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Success",
			want: "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Logger: echo.New().Logger,
			}
			if got := s.Get(); got != tt.want {
				t.Errorf("Service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
