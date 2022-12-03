package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/endpoint"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/handler/middlewares"
)

func Get(endpoint *endpoint.Endpoint) Route {
	return Route{
		Desc:     "Get",
		Group:    "Route",
		Path:     "api/v1/",
		Method:   http.MethodGet,
		Endpoint: endpoint.Get,
		Middlewares: []echo.MiddlewareFunc{
			middlewares.TimeConsumer,
		},
	}
}
