package manager

import "go-rest-framework/queryset"

type Manager interface {
	GetQuerySet() queryset.QuerySet
	Create(obj interface{})
}
