package endpoint

import (
	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/service"
)

type EndpointInterface interface {
	Get(c echo.Context) error
	Post(c echo.Context) error
}

type Endpoint struct {
	Service service.ServiceInterface
}

func Init(logger echo.Logger) *Endpoint {
	return &Endpoint{
		Service: service.Init(logger),
	}
}
