package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/endpoint"
)

type Route struct {
	Desc        string
	Group       string
	Path        string
	Method      string
	Endpoint    echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

func Init(e *echo.Echo) error {
	endpoint := endpoint.Init(e.Logger)

	routes := []Route{
		Get(endpoint),
		Post(endpoint),
	}

	for _, route := range routes {
		e.Add(route.Method, route.Path, route.Endpoint, route.Middlewares...)
	}

	return nil
}
