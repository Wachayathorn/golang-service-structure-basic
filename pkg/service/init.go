package service

import "github.com/labstack/echo/v4"

type ServiceInterface interface {
	Get() string
}

type Service struct {
	Logger echo.Logger
}

func Init(logger echo.Logger) *Service {
	return &Service{
		Logger: logger,
	}
}
