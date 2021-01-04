package resources

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-rest-framework/manager"
	"strconv"
)

var DefaultGetObject = func (m manager.Manager) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return nil, err
		}
		qs := m.GetQuerySet()
		qs.Where("id = ?", id)
		res := qs.Materialize()
		if len(res) == 0 {
			return nil, errors.New("Not found")
		}
		return res[0], nil
	}
}

type ResourceOperationBase struct {
	Manager manager.Manager
}
