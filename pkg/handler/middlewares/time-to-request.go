package middlewares

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func TimeConsumer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		if err := next(c); err != nil {
			return err
		}
		end := time.Now()
		log.Printf("Time Request id:%s from url:%s Time:%s", c.Response().Header().Get(echo.HeaderXRequestID), c.Request().RequestURI, end.Sub(start))
		return nil
	}
}
