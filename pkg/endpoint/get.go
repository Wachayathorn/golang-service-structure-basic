package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (e *Endpoint) Get(c echo.Context) error {
	res := e.Service.Get()
	return c.JSON(http.StatusOK, res)
}
