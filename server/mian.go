package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"server/config"
	"server/controller"
	"server/database"
	"server/util"
)

func main() {
	config.InitConfig(os.Args)
	database.InitClient()
	//mylog.InitMyLogger()
	e := echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	initRouter(e)
	e.Logger.Fatal(e.Start(":8088"))
}

func initRouter(e *echo.Echo) {
	controller.RouterAuth(e.Group("/auth"))
	controller.RouterUser(e.Group("/user", util.CheckJwt))
	controller.RouterSystem(e.Group("/system", util.CheckJwt))
	controller.RouterCategory(e.Group("/category", util.CheckJwt))
	controller.RouterTechnology(e.Group("/technology", util.CheckJwt))
}
