package serializers

import (
	"github.com/gin-gonic/gin"
)

type Serializer interface {
	Serialize(interface{}, *gin.Context) interface{}
	//SerializeList(interface{}, *gin.Context) ResultObj
}

