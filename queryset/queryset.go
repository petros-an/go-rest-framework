package queryset

type QuerySet interface {
	Where(query string, args ...interface{}) QuerySet
	WhereID(int) QuerySet
	Materialize() []interface{}
	Preload(column string, conditions ...interface{}) QuerySet
	Limit(int) QuerySet
	Offset(int) QuerySet
	Count() int
	Order(fieldName string, direction string) QuerySet
	Delete()
	Updates(interface{})
	FindOne() interface{}
	Exists() bool
}

