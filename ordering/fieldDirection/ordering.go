package fieldDirection

import "go-rest-framework/queryset"

type FieldDirectionOrdering struct {
	FieldName string
	Direction string
	DBFieldName string
}

func (o FieldDirectionOrdering) OrderQuerySet(qs queryset.QuerySet) {
	qs.Order(o.DBFieldName, o.Direction)
}
