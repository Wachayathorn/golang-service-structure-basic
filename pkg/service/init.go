package service

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/model"
)

type ServiceInterface interface {
	Get() string
	Post(req model.Post) (string, error)
}

type Service struct {
	Logger    echo.Logger
	Validator *validator.Validate
}

func Init(logger echo.Logger) *Service {
	return &Service{
		Logger:    logger,
		Validator: validator.New(),
	}
}
