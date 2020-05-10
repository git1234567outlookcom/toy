package controller

import (
	"fmt"
	"math"
	"server/model"
	"testing"
)

func TestName(t *testing.T) {
	var users []*model.User

	m := new(model.User)
	m.Email = "dev@dev.com"
	users = append(users, m)

	fmt.Printf("%+v", users)

	page := model.PageInfo{
		PageNum:   0,
		PageCount: 0,
		Limit:     10,
		Count:     9,
	}
	page.PageCount = int64(math.Ceil(float64(page.Count) / float64(page.Limit)))

	println(page.PageCount)
	println()
	sprintf := fmt.Sprintf("%s", "获取信息失败")
	println(sprintf)
	s := fmt.Sprintf("获取信息失败")

	println(s)
	println()
}
