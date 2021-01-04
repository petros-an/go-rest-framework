package fields

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BoolField struct {
	FieldBase
}


func (f BoolField) Deserialize(obj interface{}, c *gin.Context) (interface{}, []string, error) {
	val, err := f.FieldBase.CheckExistance(obj, c)
	if err != nil {
		return nil, nil, err
	}
	if val == nil {
		return nil, nil, nil
	}
	boolVal, ok := val.(bool)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("Invalid bool value: %s", val))
	}
	return boolVal, []string{f.InputName}, nil
}
