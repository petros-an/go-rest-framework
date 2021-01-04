package fields

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type IntField struct {
	FieldBase
	Min *int
	Max *int
}

func (f IntField) Deserialize(obj interface{}, c *gin.Context) (interface{}, []string, error) {
	val, err := f.FieldBase.CheckExistance(obj, c)
	if err != nil {
		return nil, nil, err
	}
	if val == nil {
		return nil, nil, nil
	}
	floatVal, ok := val.(float64)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("Invalid int value: %s", val))
	}
	intVal := int(floatVal)
	if f.Min != nil && intVal < *f.Min {
		return nil, nil, errors.New(fmt.Sprintf("Field minimum is %d", *f.Min))
	}
	if f.Max != nil && intVal > *f.Max {
		return nil, nil, errors.New(fmt.Sprintf("Field maximum is %d", *f.Max))
	}
	return intVal, []string{f.InputName}, nil
}
