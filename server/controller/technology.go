package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"server/model"
	"server/service"
	"server/util"
)

type TechnologyController struct {
	Service *service.TechnologyService
}

var tc TechnologyController

func RouterTechnology(g *echo.Group) {
	tc.Service = service.GetTechnologyService()
	g.POST("/", tc.Save)
	g.GET("/list", tc.FindList)
	g.DELETE("/:id", tc.Delete)
}

func (c *TechnologyController) FindList(e echo.Context) error {
	page := new(model.Page)
	if err := e.Bind(page); err != nil {
		e.Logger().Error(err)
		page.Default()
	}
	find, err := c.Service.Find(nil, bson.M{}, page.ToOptionFind())
	if err != nil {
		return Res503(e, err)
	}
	var result []*model.Technology
	for find.Next(nil) {
		m := new(model.Technology)
		err := find.Decode(m)
		if err != nil {
			return Res503(e, err)
		}
		result = append(result, m.SetId())
	}
	page.Count, err = c.Service.CountDocuments(nil, bson.M{})
	if err != nil {
		return Res503(e, err)
	}
	return Res200List(e, result, page.SetPageCount())
}

func (c *TechnologyController) Delete(e echo.Context) error {
	id := e.Param("id")
	if err := util.CheckId(id); err != nil {
		return Res400(e)
	}
	_, err := c.Service.DeleteOne(nil, bson.M{"_id": util.ToObjectId(id)})
	if err != nil {
		return Res503(e, err)
	}
	return Res200(e)
}

func (c *TechnologyController) Save(e echo.Context) error {
	m := new(model.Technology)
	if err := e.Bind(m); err != nil {
		return Res400Err(e, err)
	}
	_, err := c.Service.InsertOne(nil, m.SetObjectId)
	if err != nil {
		return Res503(e, err)
	}
	return Res200(e)
}
