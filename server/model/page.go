package model

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
)

type Page struct {
	PageNum   int64 `json:"pageNum,omitempty" desc:"第几页"`
	PageCount int64 `json:"countPage,omitempty" desc:"总页数"`
	Limit     int64 `json:"limit" desc:"每页多少条"`
	Count     int64 `json:"count,omitempty" desc:"总条数"`
}

func (p *Page) SetPageCount() *Page {
	p.PageCount = int64(math.Ceil(float64(p.Count) / float64(p.Limit)))
	return p
}
func (p *Page) Default() *Page {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	if p.PageNum <= 0 {
		p.PageNum = 1 //从第一页开始
	}
	return p
}
func (p *Page) ToOptionFind() *options.FindOptions {
	option := options.Find()
	if p.Limit <= 0 || p.Limit > 1000 {
		p.Limit = 10
	}
	option.SetLimit(p.Limit)
	if p.PageNum <= 0 {
		p.PageNum = 1 //从第一页开始
	}
	option.SetSkip(p.Limit * (p.PageNum - 1))
	return option
}
