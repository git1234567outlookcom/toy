package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"server/config"
	"server/controller"
	"server/database"
)

func main() {
	config.InitConfig(os.Args)
	database.InitClient()
	//mylog.InitMyLogger()
	e := echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	controller.RouterAuth(e.Group("/auth"))
	controller.RouterUser(e.Group("/user"))
	controller.RouterSystem(e.Group("/system"))
	controller.RouterCategory(e.Group("/category"))
	controller.RouterTechnology(e.Group("/technology"))
	//CheckJwt()
	e.Logger.Fatal(e.Start(":8088"))
}
