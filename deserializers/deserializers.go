package deserializers

import "github.com/gin-gonic/gin"

type Deserializer interface {
	Deserialize(interface{}, *gin.Context) (interface{},[]string, error)
}

type ValidatorFunction func(interface{}, *gin.Context) (interface{}, error)