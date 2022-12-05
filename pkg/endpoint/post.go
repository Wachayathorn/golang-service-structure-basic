package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/model"
)

func (e *Endpoint) Post(c echo.Context) error {
	var data model.Post
	if err := c.Bind(&data); err != nil {
		c.Logger().Errorf("Binding request body error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := e.Service.Post(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
