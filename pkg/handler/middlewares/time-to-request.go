package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
)

func TimeConsumer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		start := time.Now()
		if err := next(e); err != nil {
			return err
		}
		end := time.Now()
		e.Logger().Printf("Time Request id:%s from url:%s Time:%s", e.Response().Header().Get(echo.HeaderXRequestID), e.Request().RequestURI, end.Sub(start))
		return nil
	}
}
