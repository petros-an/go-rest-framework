package fields

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type JSONField struct {
	FieldBase
}

func (f JSONField) Deserialize(obj interface{}, c *gin.Context) (interface{}, []string, error) {
	val, err := f.FieldBase.CheckExistance(obj, c)
	if val == nil {
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, err
	}

	intVal, ok := val.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("Invalid json value: %s", val))
	}
	return intVal, []string{f.InputName}, nil
}

