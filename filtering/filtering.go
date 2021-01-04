package filtering

import (
	"github.com/gin-gonic/gin"
	"go-rest-framework/queryset"
)

type FilterCreator interface {
	GetFilters(c *gin.Context) ([]Filter, error)
}

type Filter interface {
	FilterQuerySet(qs queryset.QuerySet)
}