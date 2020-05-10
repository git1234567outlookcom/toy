package test

import (
	"encoding/json"
	"server/controller"
	"server/model"
	"server/service"
	"testing"
)

func TestSaveUser(t *testing.T) {

	user := model.User{
		Username: "123",
		Password: "123",
		Phone:    "123",
		Age:      0,
		Email:    "123@123.com",
	}
	marshal, _ := json.Marshal(user)
	body := string(marshal)

	//res, context := preparePostBody("{\"username\":\"admin\",\"password\":\"admin\",\"email\":\"admin@admin\"}")
	res, context := preparePostBody(body)
	userController := controller.UserController{Service: service.GetUserService()}
	userController.Save(context)

	t.Log(res.Code, res.Body.String())
	if res.Code != 200 {
		t.Fatal(res.Body.String())
	}

}

func TestDrop(t *testing.T) {
	initTestConfig()
	//service.GetCategoryService().Dao.Drop(nil)
	service.GetUserService().Dao.Client.Drop(nil)
}
