package queryParamFilterCreator

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/filtering"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/filtering/fieldFilter"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/filtering/operators"
	"strconv"
)

type QPFilterOption struct {
	QPName           string
	Type             string
	AssosciatedField string
	Operator         string
	Required         bool
}

type QueryParamFilterCreator struct {
	Options []QPFilterOption
}

func New(options []QPFilterOption) *QueryParamFilterCreator {
	return &QueryParamFilterCreator{Options: options}
}

func (creator *QueryParamFilterCreator) GetFilters(c *gin.Context) ([]filtering.Filter, error) {
	res := make([]filtering.Filter, 0)
	for _, opt := range creator.Options {
		qpValue, err := GetQueryParam(c, opt)
		if err != nil {
			return nil, err
		}
		if qpValue != nil {
			res = append(res, fieldFilter.FieldFilter{opt.AssosciatedField, opt.Operator, qpValue})
		}
	}
	return res, nil
}

func GetQueryParam(c *gin.Context, opt QPFilterOption) (interface{}, error){
	qpStringVal := c.Query(opt.QPName)
	if qpStringVal == "" {
		if opt.Required {
			return nil, errors.New(fmt.Sprintf("Query param %s must be provided", opt.QPName))
		} else {
			return nil, nil
		}
	}
	switch opt.Type {
	case "string":
		return qpStringVal, nil
	case "int":
		intVal, err := strconv.Atoi(qpStringVal)
		if err != nil {
			return nil, err
		}
		return intVal, nil
	case "bool":
		if qpStringVal == "true" {
			return true, nil
		} else if qpStringVal == "false" {
			return false, nil
		} else {
			return nil, errors.New(fmt.Sprintf("Invalid boolean value: %s", qpStringVal))
		}
	default:
		return nil, errors.New("Unknown query param type")
	}
}

var BeforeAfterQPFilterCreator = New([]QPFilterOption{
	{
		"before", "int", "timestamp", operators.LTE, false,
	},
	{
		"after", "int", "timestamp", operators.GTE, false,
	},
})