package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/model"
	"server/util"
	"time"
)

const PhoneOrPasswordIsNil = "用户名、密码不能为空"

type Controller struct {
	Name    string
	Model   interface{}
	Service interface{}
	Dao     interface{}
	//*CategoryController
}

var controller Controller

func Init() {
	//controller = Controller{CategoryController: new(CategoryController)}
}

func RegisterController() {
	// 统一初始化
	//controller.uc.Service = service.GetUserService()
	//controller.us = service.GetUserService()
	//service.GetRoleService()
}

func Res400(e echo.Context) error { // 客户端 请求参数
	return e.JSON(400, model.ResErr{
		Message:   util.ParameterError,
		Timestamp: time.Now().Unix(),
	})
}

func Res400Msg(e echo.Context, msg string) error { // 客户端 请求参数
	return e.JSON(200, model.ResErr{
		Message: msg,
		//Code:      400,
		Timestamp: time.Now().Unix(),
	})
}
func Res400Err(e echo.Context, err error) error { // 客户端 请求参数
	return e.JSON(200, model.ResErr{
		Error:     err,
		Message:   util.ParameterError,
		Timestamp: time.Now().Unix(),
	})
}

func Res503(context echo.Context, err error) error { // 服务端 dao 报错
	return context.JSON(200, model.ResErr{
		//Error:   err.Error(),
		Message: util.ERROR,
		//Code:      503,
		Timestamp: time.Now().Unix(),
	})
}

func Res200(e echo.Context) error {
	return e.JSON(200, model.Res{Code: 200, Timestamp: time.Now().Unix(), Message: "success"})
}

func Res200Msg(e echo.Context, msg string) error {
	return e.JSON(200, model.Res{Code: 200, Timestamp: time.Now().Unix(), Message: msg})
}
func Res200Data(e echo.Context, data interface{}) error {
	return e.JSON(http.StatusOK, model.Res{Code: 200, Timestamp: time.Now().Unix(), Message: "success", Data: data})
}
func Res200List(e echo.Context, data interface{}, page *model.Page) error {
	return e.JSON(http.StatusOK, model.Res{Data: data, PageInfo: page, Code: 200, Timestamp: time.Now().Unix(), Message: "success"})
}
