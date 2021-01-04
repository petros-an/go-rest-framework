package fields

import (
	"github.com/gin-gonic/gin"
	"go-rest-framework/deserializers"
)

type NestedField struct {
	FieldBase
	Deserializer deserializers.Deserializer
}


func (f NestedField) Deserialize(obj interface{}, c *gin.Context) (interface{}, []string, error) {
	nestedObject := obj.(map[string]interface{})[f.InputName]
	val, err := f.CheckExistance(obj, c)
	if err != nil {
		return nil, nil, err
	}
	if val == nil {
		return nil, nil, nil
	}
	if nestedObject == nil {
		return nil, nil, nil
	}
	val, _, err = f.Deserializer.Deserialize(nestedObject, c)
	if err != nil {
		return nil, nil, err
	}
	return val, []string{f.InputName}, nil
}
