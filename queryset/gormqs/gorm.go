package gormqs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/petros-an/github.com/petros-an/go-rest-framework/queryset"
	"reflect"
)

type GORMQuerySet struct {
	Model interface{}
	DB *gorm.DB
}

func (qs *GORMQuerySet) WhereID(id int) queryset.QuerySet {
	qs.DB = qs.DB.Where("id = ?", id)
	return qs
}

func (qs *GORMQuerySet) Exists() bool {
	return len(qs.Materialize()) > 0
}

func (qs *GORMQuerySet) Order(fieldName string, direction string) queryset.QuerySet {
	qs.DB = qs.DB.Order(fmt.Sprintf("%s %s", fieldName, direction))
	return qs
}

func (qs *GORMQuerySet) Limit(l int) queryset.QuerySet {
	qs.DB = qs.DB.Limit(l)
	return qs
}

func (qs *GORMQuerySet) Offset(o int) queryset.QuerySet {
	qs.DB = qs.DB.Offset(o)
	return qs
}

func (qs *GORMQuerySet) Where(query string, args ...interface{}) queryset.QuerySet {
	qs.DB = qs.DB.Where(query, args)
	return qs
}

func (qs *GORMQuerySet) Materialize() []interface{} {
	dtype := reflect.TypeOf(qs.Model)
	out := reflect.New(reflect.SliceOf(dtype)).Interface()
	x := qs.DB.Find(out)
	if x.Error != nil {
		print(x.Error.Error())
	}
	sliceOfResults := reflect.Indirect(reflect.ValueOf(out))
	len := sliceOfResults.Len()
	var res []interface{}
	for i:= 0; i < len; i++ {
		elem := sliceOfResults.Index(i).Interface()
		res = append(res, elem)
	}
	return res
}

func (qs *GORMQuerySet) Count() int {
	var count int
	qs.DB.Count(&count)
	return count
}

func (qs *GORMQuerySet) Updates(obj interface{}) {
	err := qs.DB.Updates(obj)
	print(err)
}

func (qs *GORMQuerySet) Delete() {
	qs.DB.Delete(qs.Model)
}

func (qs *GORMQuerySet) Preload(column string, conditions ...interface{}) queryset.QuerySet {
	if len(conditions) == 0 {
		qs.DB = qs.DB.Preload(column)
	} else {
		qs.DB = qs.DB.Preload(column, conditions)
	}
	return qs
}

func (qs *GORMQuerySet) FindOne() interface{} {
	items := qs.Materialize()
	if len(items) == 0 {
		return nil
	}
	return items[0]
}

func New(db *gorm.DB, model interface{}) *GORMQuerySet {
	return &GORMQuerySet{DB: db.Model(model), Model: model}
}
