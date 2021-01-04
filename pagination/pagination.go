package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/pagination/limitOffsetPagination"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/queryset"
)

type Paginator interface {
	GetPagination(c *gin.Context) (Pagination, error)
	GetResponse(data interface{}, count int) gin.H
}

type Pagination interface {
	PaginateQuerySet(set queryset.QuerySet) error
	GetElasticParams() (int, int)
}

var DefaultPaginator = limitOffsetPagination.NewPaginator(20, 0)