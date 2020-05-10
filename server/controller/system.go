package controller

import (
	"github.com/labstack/echo/v4"
	"runtime"
	"server/database"
	"server/model"
	"strconv"
)

//type SystemController2 Controller todo  有时间再玩下

type SystemController struct {
}

func (c SystemController) Get(context echo.Context) error {

	// todo 使用socket 实时获取系统信息
	m := new(model.SystemInfo)
	m.Arch = runtime.GOARCH
	m.Os = runtime.GOOS
	m.Environment = runtime.Version()
	m.Database = database.DB.Name()
	m.Cpu = strconv.Itoa(runtime.NumCPU())

	return Res200Data(context, m)
}

var sc SystemController

func RouterSystem(g *echo.Group) {
	g.GET("", sc.Get)
}
