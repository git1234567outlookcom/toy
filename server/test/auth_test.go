package test

import (
	"server/config"
	"server/controller"
	"server/database"
	"server/model"
	"server/util"
	"testing"
)

func TestLogin(t *testing.T) {
	res, context := preparePost()
	controller.Login(context)
	t.Log(res.Code, res.Body.String())
	if res.Code != 200 {
		t.Fatal(res.Body.String())
	}
}

func TestMy(t *testing.T) {
	t.Log(len("5ea8d6de1df94d1f51a06cad"))
}

func TestName(t *testing.T) {
	u := new(model.User)
	u.Phone = "15238030609"
	u.Username = "dev"
	token := util.GenerateToken(u)
	t.Log(token)
}

func TestValidateToken(t *testing.T) {

	config.InitConfig([]string{})
	database.InitClient()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODg5MzkxNTMsImlzcyI6ImRldmVsb3BlciIsInBob25lIjoiMTUyMzgwMzA2MDkiLCJ1c2VybmFtZSI6ImRldiJ9.HY3MTqfpQcGd7i8psPap_pNvVEqsx5S9bl6fIRIcWw0"
	validateToken, err := util.ValidateToken(token)
	if err != nil {
		t.Error(err)
	} else {
	}

	t.Logf("%v+", validateToken)

}
