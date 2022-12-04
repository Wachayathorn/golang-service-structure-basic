package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/endpoint"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/handler/middlewares"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/httpinfo"
)

func Get(endpoint *endpoint.Endpoint) Route {
	return Route{
		Desc:     "Get",
		Group:    "Route",
		Path:     httpinfo.Get.ApiPath,
		Method:   httpinfo.Get.HttpMethod,
		Endpoint: endpoint.Get,
		Middlewares: []echo.MiddlewareFunc{
			middlewares.TimeConsumer,
		},
	}
}
