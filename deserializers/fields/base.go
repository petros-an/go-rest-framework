package fields

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type FieldBase struct {
	InputName string
	Required bool
	Default interface{}
}

func (f FieldBase) CheckExistance(obj interface{}, c *gin.Context) (interface{}, error) {
	value := obj.(map[string]interface{})[f.InputName]
	if value == nil {
		if f.Required {
			return nil, errors.New(fmt.Sprintf("Field %s is missing", f.InputName))
		} else {
			return f.Default, nil
		}
	}
	return value, nil
}
