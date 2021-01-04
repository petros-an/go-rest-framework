package fieldFilter

import (
	"fmt"
	"go-rest-framework/queryset"
)

type FieldFilter struct {
	FieldName string
	Operator string
	Value interface{}
}

func (f FieldFilter) FilterQuerySet(qs queryset.QuerySet) {
	query := fmt.Sprintf("%s %s ? ", f.FieldName, f.Operator)
	qs.Where(query, f.Value)
}

