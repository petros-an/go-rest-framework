package list

import (
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/filtering"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/manager"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/ordering"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/pagination"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/queryset"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/resources"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/serializers"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/types"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/utils"
	"net/http"
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
		r.Orderer = ordering.DefaultOrderer
	}
}

func AddPaginator(r *ResourceList) {
	if r.Paginator == nil {
		r.Paginator = pagination.DefaultPaginator
	}
}

func AddfilterCreator(r *ResourceList) {
	if r.FilterCreator == nil {
		r.FilterCreator = filtering.DefaultFilterCreator
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


