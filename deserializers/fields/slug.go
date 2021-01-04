package fields

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-framework/manager"
)

type SlugRelatedField struct {
	FieldBase
	Manager manager.Manager
	SlugField string
}


func (f SlugRelatedField) Deserialize(obj interface{}, c *gin.Context) (interface{}, []string, error) {
	val, err := f.FieldBase.CheckExistance(obj, c)
	if err != nil {
		return nil, nil, err
	}
	if val == nil {
		return nil, nil, nil
	}
	strVal, ok := val.(string)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("Invalid slug value: %s", val))
	}
	qs := f.Manager.GetQuerySet().Where(fmt.Sprintf("%s = ?", f.SlugField), strVal)
	res := qs.Materialize()
	if len(res) == 0 {
		return nil, nil, errors.New("slug related object not found")
	}
	return res[0], []string{f.InputName}, nil
}
