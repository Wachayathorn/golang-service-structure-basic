package client

import (
	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/model"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/utils"
)

type ClientAPI interface {
	ReplaceUrl(url string)

	Get() (string, error)
	Post(req model.Post) (string, error)
}

type Client struct {
	Url   string
	Utils utils.UtilsInterface
}

func Init(e *echo.Echo) *Client {
	return &Client{
		Url:   "http://localhost:8080",
		Utils: utils.Init(e.Logger),
	}
}

func (c *Client) ReplaceUrl(url string) {
	c.Url = url
}
