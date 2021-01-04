package detail

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/manager"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/resources"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/serializers"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/types"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/utils"
)

type ResourceDetail struct {
	Serializer serializers.Serializer
	GetObject  types.GetObjectFunc
	resources.ResourceOperationBase
}

func (r *ResourceDetail) GetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		object, err := r.GetObject(c)
		if err != nil {
			utils.SendBadRequest(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": r.Serializer.Serialize(object, c),
		})
	}
}

func New(r ResourceDetail, mgr manager.Manager) types.ResourceOperation {
	r.Manager = mgr
	AddGetObject(&r)
	return &r
}

func AddGetObject(r *ResourceDetail) {
	if r.GetObject == nil {
		r.GetObject = resources.DefaultGetObject(r.Manager)
	}
}