package objectSerializer

import (
	"github.com/gin-gonic/gin"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/serializers"
	"reflect"
)

type ObectSerializer struct {
	Fields []SerializedField
}

type SerializedField struct {
	DisplayedName string
	Serializer    serializers.Serializer
}

func (s *ObectSerializer) Serialize(obj interface{}, c *gin.Context) interface{} {
	res := map[string]interface{}{}
	for _, f := range s.Fields {
		res[f.DisplayedName] = f.Serializer.Serialize(obj, c)
	}
	return res
}

type FieldSerializer struct {
	FieldName string
}

func (s *FieldSerializer) Serialize(obj interface{}, c *gin.Context) interface{} {
	originalElement := reflect.ValueOf(obj).Elem()
	return originalElement.FieldByName(s.FieldName).Interface()
}

type NestedSerializer struct {
	Serializer serializers.Serializer
	FieldName  string
}

func (s *NestedSerializer) Serialize(obj interface{}, c *gin.Context) interface{} {
	val := reflect.ValueOf(obj).Elem().FieldByName(s.FieldName).Addr().Interface()
	return s.Serializer.Serialize(val, c)
}

type FunctionSerializer struct {
	Function func(interface{}, *gin.Context) interface{}
}

func (s *FunctionSerializer) Serialize(obj interface{}, c *gin.Context) interface{} {
	return s.Function(obj, c)
}

func Field(displayedName, fieldName string) SerializedField {
	return SerializedField{
		DisplayedName: displayedName, Serializer: &FieldSerializer{FieldName: fieldName},
	}
}

func Nested(displayedName string, fieldName string, serializer serializers.Serializer) SerializedField {
	return SerializedField{
		DisplayedName: displayedName, Serializer: &NestedSerializer{Serializer: serializer, FieldName: fieldName},
	}
}

func Func(displayedName string, function func(interface{}, *gin.Context) interface{}) SerializedField {
	return SerializedField{
		DisplayedName: displayedName, Serializer: &FunctionSerializer{Function: function},
	}
}

