package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"server/model"
	"server/service"
	"server/util"
)

/*
这里使用了mvc的模式来尝试
controller：接收请求、校验参数
service：处理业务逻辑
dao：操作mongo

这些都不是一成不变，或者说不需要一开始都要写成这样，
是需要一步步迭代、演化
才能使代码和业务更加契合
更像是水到渠成的关系

如果刻意的去套用这个mvc
反而有点邯郸学步、东施效颦
在开发的过程中产生一些不必要的羁绊
对业务的发展不是那么的友好

这个时候就需要程序员展现对代码的把控、扩展的能力
尽量的使代码具有可重入性
对下个阶段业务的演变和代码迭代有一定的预见性
使代码具有很好的扩展性
//20200507
*/
type UserController struct {
	Service *service.UserService
}

var uc UserController

func RouterUser(g *echo.Group) {
	uc.Service = service.GetUserService()
	g.POST("", uc.Save)
	g.DELETE("/:id", uc.Delete)
	g.GET("/list", uc.FindList)
	//g.PUT("/update", uc.Update) 暂时先不修改用户
}

func (c *UserController) Delete(e echo.Context) error {
	user := new(model.User)
	user.Id = e.Param("id")
	if user.Id == "" {
		return Res400Err(e, errors.New(util.IdIsNil))
	}
	err := c.Service.DeleteById(user.SetObjectId())
	if err != nil {
		return Res503(e, err)
	}
	return Res200(e)

}
func (c *UserController) Save(e echo.Context) error {
	u := new(model.User)
	if err := e.Bind(u); err != nil {
		return Res400Err(e, err)
	}
	_, err := c.Service.Save(u) // 第一种 在每个group 初始化*controller
	//_, err := service.GetUserService().Save(u)  	// 第二种直接调用 service
	//_, err := controller.us.Save(u) 				// 第三种 统一 初始化*controller
	if err != nil {
		e.Logger().Error(err)
		return Res503(e, err)
	}
	return Res200(e)
}

func (c *UserController) FindList(context echo.Context) error {
	page := new(model.Page)
	if err := context.Bind(page); err != nil {
		return Res400Err(context, err)
	}
	list, err := c.Service.FindList(page.Default())
	if err != nil {
		Res503(context, err)
	}
	return Res200List(context, list, page)
}

func (c *UserController) Update(e echo.Context) error {
	u := new(model.User)
	if err := e.Bind(u); err != nil {
		return Res400Err(e, err)
	}
	if u.Id == "" {
		return Res400(e)
	}
	update, err := c.Service.Update(u.SetObjectId())
	if err != nil {
		return Res503(e, err)
	}
	return Res200Data(e, update.SetId())
}
