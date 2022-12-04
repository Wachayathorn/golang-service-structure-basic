package utils

import "github.com/labstack/echo/v4"

type UtilsInterface interface {
	DoRequest(url, httpMethod, path string, data interface{}) ([]byte, error)
}

type Utils struct {
	Logger echo.Logger
}

func Init(logger echo.Logger) *Utils {
	return &Utils{
		Logger: logger,
	}
}
