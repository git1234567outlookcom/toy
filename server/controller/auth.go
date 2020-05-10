package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/model"
	"server/service"
	"server/util"
	"strconv"
	"time"
)

type Auth struct {
}

func RouterAuth(g *echo.Group) {
	g.POST("/login", Login)
	g.POST("/register", Register)
	g.POST("/logout", Logout)
	g.GET("/captcha", GetCaptcha)
}

func Logout(c echo.Context) error {
	return Res200(c)
}

func Login(c echo.Context) error {
	m := new(model.User)
	c.Bind(m)
	if m.Phone == "" || m.Password == "" {
		return Res400Msg(c, PhoneOrPasswordIsNil)
	}
	user, err := service.GetUserService().FindPhoneOrEmail(m)
	if err != nil {
		c.Logger().Error(err)
		return Res400Msg(c, "账号不存在")
	}
	if user.Password != m.Password {
		return Res400Msg(c, "密码不正确")
	}
	token := util.GenerateToken(user)
	s := struct {
		Token string `json:"token"`
	}{Token: token}
	return Res200Data(c, s)
}

func Register(c echo.Context) error {
	m := new(model.User)
	c.Bind(m)
	// todo 在struct里面校验
	if m.Phone == "" { //todo 手机号位数
		return Res400Err(c, errors.New("手机号不能为空"))
	}

	//if m.Username == "" {
	//	return Res400Err(c, errors.New("名称不能为空"))
	//}
	if m.Password == "" {
		return Res400Err(c, errors.New("密码不能为空"))
	}
	//if m.Email == "" {
	//	return Res400Err(c, errors.New("邮箱不能为空"))
	//}

	//_, err := service.GetUserService().Dao.FindOneByFilter(bson.M{"email": m.Email})
	//if err != mongo.ErrNoDocuments {
	//	return Res200Msg(c, "该邮箱已经注册过")
	//}

	_, err := service.GetUserService().Dao.FindOneByFilter(bson.M{"phone": m.Phone})
	if err != mongo.ErrNoDocuments {
		return Res200Msg(c, "该手机号已经注册过")
	}

	_, err = service.GetUserService().Save(m)
	if err != nil {
		return Res503(c, err)
	}
	return Res200(c)
}

func GetCaptcha(c echo.Context) error {
	//模拟短信验证码
	s := struct {
		CaptchaCode string `json:"captchaCode"`
	}{CaptchaCode: strconv.FormatInt(time.Now().Unix(), 10)[6:]}
	return Res200Data(c, s)
}
