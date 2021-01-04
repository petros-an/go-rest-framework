package fieldDirection

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/ordering"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/utils"
)

type FieldDirectionOrderer struct {
	DefaultFieldName string
	DefaultDirection string
	AllowedFieldNames []string
	DBMappings map[string]string
}


func New(defaultFieldName string, defaultDirection string, allowed []string, mappings map[string]string) *FieldDirectionOrderer {
	return &FieldDirectionOrderer{
		DefaultFieldName: defaultFieldName,
		DefaultDirection: defaultDirection,
		AllowedFieldNames: allowed,
		DBMappings: mappings,
	}
}

func (o *FieldDirectionOrderer) GetOrdering(c *gin.Context) (ordering.Ordering, error) {
	orderingParam := c.Query("ordering")
	if len(orderingParam) < 2 {
		return FieldDirectionOrdering{
			FieldName: o.DefaultFieldName,
			Direction: o.DefaultDirection,
			DBFieldName: o.DBMappings[o.DefaultFieldName],
		}, nil
	}
	directionParam := orderingParam[:1]
	fieldNameParam := orderingParam[1:]
	var direction string
	if directionParam == "+" {
		direction = ordering.Asc
	} else if directionParam == "-" {
		direction = ordering.Desc
	} else {
		return nil, errors.New(fmt.Sprintf("Ordering direction not understood: %s", direction))
	}
	if !utils.ArrayContains(o.AllowedFieldNames, fieldNameParam) {
		return nil, errors.New(fmt.Sprintf("Unknown field: %s", fieldNameParam))
	}
	return FieldDirectionOrdering{
		FieldName: fieldNameParam,
		Direction: direction,
		DBFieldName: o.DBMappings[fieldNameParam],
	}, nil
}

