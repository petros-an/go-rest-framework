package create

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/deserializers"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/manager"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/resources"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/types"
)

type ResourceCreate struct {
	Model interface{}
	Deserializer  deserializers.Deserializer
	PerformCreate types.PerformCreateFunc
	resources.ResourceOperationBase
}

func (r *ResourceCreate) GetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		bts, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var inp map[string]interface{}
		err = json.Unmarshal(bts, &inp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		obj, _, desErr := r.Deserializer.Deserialize(inp, c)
		if desErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": desErr.Error(),
			})
			return
		}

		err = r.PerformCreate(obj)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Status(http.StatusCreated)
	}
}

func New(r ResourceCreate, mgr manager.Manager) types.ResourceOperation {
	r.Manager = mgr
	AddPerformCreate(&r)
	return &r
}

func AddPerformCreate(r *ResourceCreate) {
	if r.PerformCreate == nil {
		r.PerformCreate = func(obj interface{}) error {
			r.Manager.Create(obj)
			return nil
		}
	}
}