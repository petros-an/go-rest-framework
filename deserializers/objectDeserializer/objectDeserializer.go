package objectDeserializer

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-framework/deserializers"
	f "go-rest-framework/deserializers/fields"
	"go-rest-framework/manager"
)

type Field struct {
	ModelField   string
	Deserializer deserializers.Deserializer
}

type ObjectDeserializer struct {
	Fields []Field
	Validators map[string]deserializers.ValidatorFunction
}

type ValidatorMap map[string]func(interface{}, *gin.Context) error

func (d *ObjectDeserializer) Deserialize(obj interface{}, c *gin.Context) (interface{}, []string, error) {
	resultMap := map[string]interface{}{}
	consumedFields := []string{}
	for _, field := range d.Fields {
		val, fields, err := field.Deserializer.Deserialize(obj, c)
		if err != nil {
			return nil, nil, err
		}
		if val == nil {
			continue
		}
		validator := d.Validators[field.ModelField]
		if validator != nil {
			newVal, err := validator(obj, c)
			if err != nil {
				return nil, nil, err
			}
			val = newVal
		}
		resultMap[field.ModelField] = val
		consumedFields = append(consumedFields, fields...)
	}
	allowedFields := map[string]bool{}
	for _, f := range consumedFields{
		allowedFields[f] = true
	}
	for k,_ := range obj.(map[string]interface{}) {
		if _, exists := allowedFields[k]; !exists {
			return nil, nil, errors.New(fmt.Sprintf("Unknown field: %s", k))
		}
	}
	return resultMap, consumedFields, nil
}

func Bool(inpName, modelField string) Field {
	return Field{ModelField: modelField, Deserializer: f.BoolField{FieldBase: f.FieldBase{InputName: inpName}}}
}

func String(inpName, modelField string) Field {
	return Field{ModelField: modelField, Deserializer: f.StringField{FieldBase: f.FieldBase{InputName: inpName}}}
}

func Int(inpName, modelField string) Field {
	return Field{ModelField: modelField, Deserializer: f.IntField{FieldBase: f.FieldBase{InputName: inpName}}}
}

func JSON(inpName, modelField string) Field {
	return Field{ModelField: modelField, Deserializer: f.JSONField{FieldBase: f.FieldBase{InputName: inpName}}}
}

func Nested(inpName, modelField string, des deserializers.Deserializer) Field {
	return Field{ModelField: modelField, Deserializer: f.NestedField{FieldBase: f.FieldBase{InputName: inpName}, Deserializer: des}}
}

func SlugRelated(inpName, modelField string, manager manager.Manager, slug string) Field {
	return Field{ModelField: modelField, Deserializer: f.SlugRelatedField{FieldBase: f.FieldBase{InputName: inpName}, Manager: manager, SlugField: slug}}
}

func PrimaryKeyRelated(inpName, modelField string, manager manager.Manager) Field {
	return Field{ModelField: modelField, Deserializer: f.PrimaryKeyRelatedField{FieldBase: f.FieldBase{InputName: inpName}, Manager: manager}}
}
