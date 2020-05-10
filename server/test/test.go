package test

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http/httptest"
	"server/config"
	"server/database"
	"server/model"
	"server/util"
	"strings"
)

func preparePost() (*httptest.ResponseRecorder, echo.Context) {
	return prepare(echo.POST, nil)
}
func preparePostBody(body string) (*httptest.ResponseRecorder, echo.Context) {

	return prepare(echo.POST, strings.NewReader(body))
}

func prepare(method string, r io.Reader) (*httptest.ResponseRecorder, echo.Context) {
	initTestConfig()
	e := echo.New()
	//payload := strings.NewReader("{\"name\":\"Joe\",\"email\":\"joe@labstack\"}")

	//req := httptest.NewRequest(method, "/", httputil.NewChunkedReader(r))
	req := httptest.NewRequest(method, "/", r)
	//req.Header.Set("x-authorization", getTestHeaderToken())
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	return res, context
}

func initTestConfig() {
	config.InitConfig([]string{})
	database.InitClient()
}

func getTestHeaderToken() string {
	return "Bearer " + util.GenerateToken(&model.User{
		Username: "admin",
		Email:    "dev@dev.com",
	})
}
