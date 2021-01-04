package detail

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-rest-framework/manager"
	"go-rest-framework/resources"
	"go-rest-framework/serializers"
	"go-rest-framework/types"
	"go-rest-framework/utils"
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