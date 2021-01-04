package pagination

import (
	"github.com/gin-gonic/gin"
	"go-rest-framework/queryset"
)

type Paginator interface {
	GetPagination(c *gin.Context) (Pagination, error)
	GetResponse(data interface{}, count int) gin.H
}

type Pagination interface {
	PaginateQuerySet(set queryset.QuerySet) error
	GetElasticParams() (int, int)
}
