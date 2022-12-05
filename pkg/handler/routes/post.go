package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/endpoint"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/handler/middlewares"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/httpinfo"
)

func Post(endpoint *endpoint.Endpoint) Route {
	return Route{
		Desc:     "Post",
		Group:    "Route",
		Path:     httpinfo.Post.ApiPath,
		Method:   httpinfo.Post.HttpMethod,
		Endpoint: endpoint.Post,
		Middlewares: []echo.MiddlewareFunc{
			middlewares.TimeConsumer,
		},
	}
}
