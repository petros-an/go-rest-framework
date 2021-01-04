package gormManager

import (
	"phoenixserver/src/models/db"
	"go-rest-framework/manager"
	"go-rest-framework/queryset"
	"go-rest-framework/queryset/gormqs"
)

type GormV1Manager struct {
	Model interface{}
}

func (g *GormV1Manager) GetQuerySet() queryset.QuerySet {
	return gormqs.New(db.GetDB(), g.Model)
}

func (g *GormV1Manager) Create(obj interface{}) {
	db.GetDB().Create(obj)
}

func New(model interface{}) manager.Manager {
	return &GormV1Manager{
		Model: model,
	}
}