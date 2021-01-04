package manager

import "github.com/petros-an/github.com/petros-an/go-rest-framework/queryset"

type Manager interface {
	GetQuerySet() queryset.QuerySet
	Create(obj interface{})
}
