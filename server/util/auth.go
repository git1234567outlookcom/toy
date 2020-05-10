package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"server/dao"
	"server/model"
	"time"
)

const TOKEN_KEY = "liuLIU"
const TOKEN_EXPIRATION_TIME int64 = 10 // token有效期 1分钟

func GenerateToken(u *model.User) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(TOKEN_EXPIRATION_TIME)).Unix()
	claims["iss"] = "developer"
	claims["username"] = u.Username
	//claims["email"] = u.Email
	claims["phone"] = u.Phone
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sign, _ := token.SignedString([]byte(TOKEN_KEY))
	return sign
}

func ValidateToken(tokenString string) (*model.User, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(TOKEN_KEY), nil
	})
	if err != nil {
		log.Error("校验token:", err)
		return nil, errors.New("请重新登录")
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	phone := claims["phone"].(string)
	one := dao.GetUserDao().Client.FindOne(nil, bson.M{"phone": phone})
	if one.Err() != nil {
		return nil, errors.New("非法token")
	}
	m := new(model.User)
	one.Decode(m)
	return m, nil
}

func CheckJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenRaw := ctx.Request().Header.Get(echo.HeaderAuthorization) // header 查找token
		if tokenRaw == "" {
			ctx.JSON(400, model.ResErr{
				Message:   "请重新登陆",
				Code:      400,
				Timestamp: time.Now().Unix(),
			})
			return nil
		}
		tokenRaw = tokenRaw[7:] // Bearer token len("Bearer ")==7
		user, err := ValidateToken(tokenRaw)
		if err != nil {
			ctx.Logger().Error(err)
			ctx.JSON(400, model.ResErr{
				Message:   "请重新登陆",
				Code:      400,
				Timestamp: time.Now().Unix(),
			})
			return nil
		}
		ctx.Set("uid", user.Id)
		// 自定义头
		return next(ctx)
	}
}
