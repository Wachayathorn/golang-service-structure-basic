package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/config"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/handler/middlewares"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/handler/routes"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/utils"
)

func main() {
	e := echo.New()

	env := flag.String("env", "local", "Environment")
	c := config.Config{}

	port, err := c.Load(*env)
	if err != nil {
		e.Logger.Fatalf("Load config error:%s", err.Error())
	}

	routes.Init(e)
	e.Use(middlewares.TimeConsumer)
	e.Logger.Infof("Server start port:%s", port)

	if err := e.Start(":" + port); err != nil {
		e.Logger.Fatalf("Start service error%s", err.Error())
	}

	utils.Init(&e.Logger)
}
