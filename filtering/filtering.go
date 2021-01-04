package filtering

import (
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/filtering/queryParamFilterCreator"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/queryset"
)

type FilterCreator interface {
	GetFilters(c *gin.Context) ([]Filter, error)
}

type Filter interface {
	FilterQuerySet(qs queryset.QuerySet)
}

var DefaultFilterCreator = queryParamFilterCreator.New([]queryParamFilterCreator.QPFilterOption{})

