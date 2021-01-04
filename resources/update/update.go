package update

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"go-rest-framework/deserializers"
	"go-rest-framework/manager"
	"go-rest-framework/resources"
	"go-rest-framework/types"
	"reflect"
)

type ResourceUpdate struct {
	GetObject     types.GetObjectFunc
	Deserializer  deserializers.Deserializer
	PerformUpdate types.PerformUpdateFunc
	resources.ResourceOperationBase
}

func (r *ResourceUpdate) GetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		// read body
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

		// deserialize
		updObj, _, err := r.Deserializer.Deserialize(inp, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// get object to update
		obj, err := r.GetObject(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// perform update
		err = r.PerformUpdate(obj, updObj.(map[string]interface{}))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}

func New(r ResourceUpdate, mgr manager.Manager) types.ResourceOperation {
	r.Manager = mgr
	AddPerformUpdate(&r)
	AddGetObject(&r)
	return &r
}

func AddPerformUpdate(r *ResourceUpdate) {
	if r.PerformUpdate == nil {
		r.PerformUpdate = func(obj interface{}, updobj map[string]interface{}) error {
			id := reflect.ValueOf(obj).Elem().FieldByName("ID").Interface().(uint)
			qs := r.Manager.GetQuerySet()
			qs.Where("id = ?", id)
			qs.Updates(updobj)
			return nil
		}
	}
}

func AddGetObject(r *ResourceUpdate) {
	if r.GetObject == nil {
		r.GetObject = resources.DefaultGetObject(r.Manager)
	}
}