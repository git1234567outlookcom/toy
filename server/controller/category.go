package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"server/model"
	"server/service"
	"server/util"
)

type CategoryController struct {
	Service *service.CategoryService
}

var cc CategoryController

func RouterCategory(g *echo.Group) {
	cc.Service = service.GetCategoryService()
	g.POST("", cc.Save) //成了两层结构 view-service/dao
	g.PUT("", cc.Put)
	g.GET("/list", cc.FindList) //单层，提升效率，迭代时可以进行优化
	g.DELETE("/:id", cc.Delete)
}

func (c *CategoryController) FindList(e echo.Context) error {
	page := new(model.Page)
	if err := e.Bind(page); err != nil {
		e.Logger().Error(err)
		page.Default()
	}
	find, err := c.Service.Dao.Find(nil, bson.M{}, page.ToOptionFind())
	if err != nil {
		return Res503(e, err)
	}
	var result []*model.Category
	for find.Next(nil) {
		m := new(model.Category)
		err := find.Decode(m)
		if err != nil {
			return Res503(e, err)
		}
		result = append(result, m.SetId())
	}
	page.Count, err = c.Service.Dao.CountDocuments(nil, bson.M{})
	if err != nil {
		return Res503(e, err)
	}
	return Res200List(e, result, page.SetPageCount())
}

func (c *CategoryController) Delete(e echo.Context) error {
	id := e.Param("id")
	if err := util.CheckId(id); err != nil {
		return Res400(e)
	}
	_, err := c.Service.Dao.DeleteOne(nil, bson.M{"_id": util.ToObjectId(id)})
	if err != nil {
		return Res503(e, err)
	}
	return Res200(e)
}

func (c *CategoryController) Save(e echo.Context) error {
	m := new(model.Category)
	if err := e.Bind(m); err != nil {
		return Res400Err(e, err)
	}
	_, err := c.Service.Dao.InsertOne(nil, m.SetObjectId())
	if err != nil {
		return Res503(e, err)
	}
	return Res200(e)
}

func (c *CategoryController) Put(e echo.Context) error {
	m := new(model.Category)
	if err := e.Bind(m); err != nil {
		return Res400Err(e, err)
	}
	if m.Id == "" {
		return Res400Err(e, errors.New("id不能为空"))
	}
	replace := c.Service.Dao.FindOneAndReplace(nil, bson.M{"_id": m.SetObjectId().ObjectId}, m)

	if replace.Err() != nil {
		return Res503(e, replace.Err())
	}
	return Res200(e)
}
