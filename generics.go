package go_rest_framework

import (
	"go-rest-framework/filtering/queryParamFilterCreator"
	"go-rest-framework/ordering"
	"go-rest-framework/ordering/fieldDirection"
	"go-rest-framework/pagination/limitOffsetPagination"
)

var DefaultPaginator = limitOffsetPagination.NewPaginator(20, 0)
var DefaultOrderer = &fieldDirection.FieldDirectionOrderer{
	DefaultFieldName:  "id",
	DefaultDirection:  ordering.Asc,
	AllowedFieldNames: []string{"id", "createdAt"},
	DBMappings: map[string]string{
		"id": "id",
		"createdAt": "created_at",
	},
}
var DefaultFilterCreator = queryParamFilterCreator.New([]queryParamFilterCreator.QPFilterOption{})


