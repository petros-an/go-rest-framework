package delete

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/manager"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/resources"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/types"
	"reflect"
)

type ResourceDelete struct {
	resources.ResourceOperationBase
	GetObject     types.GetObjectFunc
	PerformDelete types.PerformDeleteFunc
}

func (r *ResourceDelete) GetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		obj, err := r.GetObject(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = r.PerformDelete(obj)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func New(r ResourceDelete, mgr manager.Manager) types.ResourceOperation {
	r.Manager = mgr
	AddPerformDelete(&r)
	AddGetObject(&r)
	return &r
}

func AddGetObject(r *ResourceDelete) {
	if r.GetObject == nil {
		r.GetObject = resources.DefaultGetObject(r.Manager)
	}
}

func AddPerformDelete(r *ResourceDelete) {
	r.PerformDelete = func(obj interface{}) error {
		id := reflect.ValueOf(obj).Elem().FieldByName("ID").Interface().(uint)
		qs := r.Manager.GetQuerySet()
		qs.Where("id = ?", id)
		qs.Delete()
		return nil
	}
}
