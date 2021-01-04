package limitOffsetPagination

import (
	"go-rest-framework/pagination"
	"go-rest-framework/queryset"
)

type LimitOffsetPagination struct {
	Limit int
	Offset int
}

func NewPaginator(defLim int, defOff int) pagination.Paginator {
	pg := LimitOffsetPaginator{
		DefaultLimit:  defLim,
		DefaultOffset: defOff,
	}
	return &pg
}

func (p LimitOffsetPagination) PaginateQuerySet(qs queryset.QuerySet) error {
	qs.Limit(p.Limit)
	qs.Offset(p.Offset)
	return nil
}

func (p LimitOffsetPagination) GetElasticParams() (int, int) {
	return p.Offset, p.Limit
}
