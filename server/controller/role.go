package controller

import (
	"github.com/labstack/echo/v4"
	"runtime"
	"server/database"
	"server/model"
	"strconv"
)

type RoleController struct {
}

func (c RoleController) Get(context echo.Context) error {

	m := new(model.SystemInfo)
	m.Arch = runtime.GOARCH
	m.Os = runtime.GOOS
	m.Environment = runtime.Version()
	m.Database = database.DB.Name()
	m.Cpu = strconv.Itoa(runtime.NumCPU())

	return Res200Data(context, m)
}

func RouterRole(g *echo.Group) {
	g.GET("", sc.Get)
}
