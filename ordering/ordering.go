package ordering

import (
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/ordering/fieldDirection"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/queryset"
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

var DefaultOrderer = &fieldDirection.FieldDirectionOrderer{
	DefaultFieldName:  "id",
	DefaultDirection:  Asc,
	AllowedFieldNames: []string{"id", "createdAt"},
	DBMappings: map[string]string{
		"id": "id",
		"createdAt": "created_at",
	},
}