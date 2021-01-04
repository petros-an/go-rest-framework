package fields

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/manager"
)

type PrimaryKeyRelatedField struct {
	FieldBase
	Manager manager.Manager
}


func (f PrimaryKeyRelatedField) Deserialize(obj interface{}, c *gin.Context) (interface{}, []string, error) {
	val, err := f.FieldBase.CheckExistance(obj, c)
	if err != nil {
		return nil, nil, err
	}
	if val == nil {
		return nil, nil, nil
	}
	floatVal, ok := val.(float64)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("Invalid pk value: %d", val))
	}
	intVal := int(floatVal)
	qs := f.Manager.GetQuerySet()
	qs.Where("id = ?", intVal)
	res := qs.Materialize()
	if len(res) == 0 {
		return nil, nil, errors.New("slug related object not found")
	}
	return res[0], []string{f.InputName}, nil
}
