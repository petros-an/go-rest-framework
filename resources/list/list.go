package list

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phoenixserver/src/rest"
	"go-rest-framework/filtering"
	"go-rest-framework/manager"
	"go-rest-framework/ordering"
	"go-rest-framework/pagination"
	"go-rest-framework/queryset"
	"go-rest-framework/resources"
	"go-rest-framework/serializers"
	"go-rest-framework/types"
	"go-rest-framework/utils"
)

type ResourceList struct {
	Serializer    serializers.Serializer
	GetQuerySet   types.GetQuerysetFunc
	Paginator     pagination.Paginator
	Orderer       ordering.Orderer
	FilterCreator filtering.FilterCreator
	resources.ResourceOperationBase
}

func (r *ResourceList) GetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		qs, err := r.GetQuerySet(c)
		if utils.SendBadRequest(c, err) != nil {return}

		filters, err := r.FilterCreator.GetFilters(c)
		if utils.SendBadRequest(c, err) != nil {return}
		for _, f := range filters {
			f.FilterQuerySet(qs)
		}

		count := qs.Count()

		pagination, err := r.Paginator.GetPagination(c)
		if utils.SendBadRequest(c, err) != nil {return}

		err = pagination.PaginateQuerySet(qs)
		if utils.SendBadRequest(c, err) != nil {return}

		ordering, err := r.Orderer.GetOrdering(c)
		if utils.SendBadRequest(c, err) != nil {return}
		ordering.OrderQuerySet(qs)

		objects := qs.Materialize()
		data := SerializeObjects(r.Serializer, objects, c)

		c.JSON(http.StatusOK, r.Paginator.GetResponse(data, count))
	}
}

func AddOrderer(r *ResourceList) {
	if r.Orderer == nil {
		r.Orderer = rest.DefaultOrderer
	}
}

func AddPaginator(r *ResourceList) {
	if r.Paginator == nil {
		r.Paginator = rest.DefaultPaginator
	}
}

func AddfilterCreator(r *ResourceList) {
	if r.FilterCreator == nil {
		r.FilterCreator = rest.DefaultFilterCreator
	}
}

func AddQSFunction(r *ResourceList) {
	if r.GetQuerySet == nil {
		r.GetQuerySet = func(c *gin.Context) (queryset.QuerySet, error) {
			return r.Manager.GetQuerySet(), nil
		}
	}
}

func New(r ResourceList, mgr manager.Manager) types.ResourceOperation {
	r.Manager = mgr
	AddOrderer(&r)
	AddPaginator(&r)
	AddfilterCreator(&r)
	AddQSFunction(&r)
	return &r
}

func SerializeObjects(serializer serializers.Serializer, objects []interface{}, c *gin.Context) []interface{} {
	data := []interface{}{}
	for _, obj := range objects {
		data = append(data, serializer.Serialize(obj, c))
	}
	return data
}