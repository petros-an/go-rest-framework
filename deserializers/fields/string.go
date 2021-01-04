package fields

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type StringField struct {
	FieldBase
	MinLength *int
	MaxLength *int
	Blank *bool
}

func (f StringField) Deserialize(obj interface{}, c *gin.Context) (interface{}, []string, error) {
 	val, err := f.FieldBase.CheckExistance(obj, c)
	if err != nil {
		return nil, nil, err
	}
	if val == nil {
		return nil, nil, nil
	}
	strVal, ok := val.(string)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("Invalid string value: %s", val))
	}
	if f.MinLength != nil && len(strVal) < *f.MinLength {
		return nil, nil, errors.New(fmt.Sprintf("Length minimum is %d", *f.MinLength))
	}
	if f.MaxLength != nil && len(strVal) > *f.MaxLength {
		return nil, nil, errors.New(fmt.Sprintf("Length maximum is %d", *f.MaxLength))
	}
	return strVal, []string{f.InputName}, nil
}
