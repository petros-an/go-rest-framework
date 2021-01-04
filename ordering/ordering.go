package ordering

import (
	"github.com/gin-gonic/gin"
	"go-rest-framework/queryset"
)

const (
	Asc = "asc"
	Desc = "desc"
)

type Ordering interface {
	OrderQuerySet(qs queryset.QuerySet)
}

type Orderer interface {
	GetOrdering(c *gin.Context) (Ordering, error)
}
