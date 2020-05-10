package util

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

func GetBody(c echo.Context, m interface{}) {
	body, _ := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	json.Unmarshal(body, m)
}
