package types

import (
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/queryset"
)

type ResourceOperation interface {
	GetEndpoint() gin.HandlerFunc
}

type ResultObj interface{}
type ResultList []ResultObj

type GetQuerysetFunc func(*gin.Context) (queryset.QuerySet, error)
type GetObjectFunc func(*gin.Context) (interface{}, error)
type PerformDeleteFunc func(obj interface{}) error
type PerformCreateFunc func(obj interface{}) error
type PerformUpdateFunc func(obj interface{}, updobj map[string]interface{}) error


