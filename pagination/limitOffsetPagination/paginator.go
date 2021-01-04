package limitOffsetPagination

import (
	"github.com/gin-gonic/gin"
	"go-rest-framework/pagination"
	"strconv"
)

type LimitOffsetPaginator struct {
	DefaultLimit int
	DefaultOffset int
}

func (p *LimitOffsetPaginator) Test() {}

func (p *LimitOffsetPaginator) GetPagination(c *gin.Context) (pagination.Pagination, error) {
	lim, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		lim = p.DefaultLimit
	}
	off, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		off = p.DefaultOffset
	}
	return LimitOffsetPagination{
		Limit:  lim,
		Offset: off,
	}, nil
}

func (p *LimitOffsetPaginator) GetResponse(data interface{}, count int) gin.H {
	return gin.H {
		"data": data,
		"count": count,
	}
}


