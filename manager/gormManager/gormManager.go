package gormManager

import (
	"github.com/petros-an/github.com/petros-an/go-rest-framework/manager"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/queryset"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/queryset/gormqs"
	"github.com/jinzhu/gorm"
)

type GormV1Manager struct {
	Model interface{}
	GetDB func() *gorm.DB
}

func (g *GormV1Manager) GetQuerySet() queryset.QuerySet {
	return gormqs.New(g.GetDB(), g.Model)
}

func (g *GormV1Manager) Create(obj interface{}) {
	g.GetDB().Create(obj)
}

func New(model interface{}, getdb func()*gorm.DB) manager.Manager {
	return &GormV1Manager{
		Model: model,
	}
}