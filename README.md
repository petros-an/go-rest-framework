# Django in Golang

This is a demonstration of how a Django-like REST framework can be built in Go, on top of Gin and GORM.
As in Django and DRF, the concept is to build a 'bridge' between the ORM and HTTP endpoints which provides
a way to abstract CRUD operations. Thus, I have created the following:

- The QuerySet interface, which abstracts away the ORM (and a GORM implementation) and acts as a container of
the database query that a REST resource uses
- Managers, which match a REST resources to QuerySets and provide Create methods
- Resource Operations (corresponding to Django APIViews) that abstract away the List, Retrieve, Create, Update, Delete operations
- Serializers and Deserializers (I thought it more convenient to seperate the two, as opposed to Django, where a Serializer performs both functions)
- Pagination, Filtering and Ordering

To imitate Django's Python magic which uses metaclasses and other dynamic features, we have to use 
modules, interfaces and composition instead of inheritance, metaclasses, and dynamic attributes.
For example, a Serializer in django is defined by a class with class attributes representing fields:

```

class LocationSerializer(ModelSerializer):
  name = serializers.CharField()
  ...
  
```
In Go, we can use composition:
```
var LocationSerializer = ms.ObectSerializer {
	Fields: *an array of fields*
}
    
```


A REST API with CRUD operations on a Location resource can be created as follows:

1) Gorm model definition

```
type Location struct {
	gorm.Model
	Name string
	City string 
	Country string 
	Description string
}
```

2) A Manager

```
var LocationManager = gormManager.New(&Location{})
```

3) A Serializer for the List & Detail operations
```
var LocationSerializer = ms.ObectSerializer{
	Fields: []ms.SerializedField{
		ms.Field("id", "ID"),
		ms.Field("description", "Description"),
		ms.Field("createdAt", "CreatedAt"),
		ms.Field("country", "Country"),
		ms.Field("city", "City"),
		ms.Field("name", "Name"),
	},
}
```

4) A Deserializer for the Create & Update operations
```
var LocationDeserializer = objectDeserializer.ObjectDeserializer{
	Fields: []ds.Field{
		ds.String("city", "City"),
		ds.String("country", "Country"),
		ds.String("name", "Name"),
		ds.String("description", "Description"),
	},
}
```

5) Define the operations

```
	ResourceDetail = detail.New(
		detail.ResourceDetail {
			Serializer: &LocationSerializer,
		},
		location.LocationManager,
	)
  
  /*
    Let's add url param filtering functionality to the List operation
  */
  
  ResourceList = list.New(
		list.ResourceList {
			Serializer: &LocationSerializer,
			FilterCreator: queryParamFilterCreator.New(
				[]queryParamFilterCreator.QPFilterOption{
					{
						"city", "string", "city", operators.EQ, false,
					},
				},
			),
		},
		location.LocationManager,
	)
  
  /*
    And similarly for the other operations
  */
```

6) Create the endpoints
```
  /* engine is a GIN *Engine object. GetEndpoint() is the counterpart to Django's APIView.as_view() */
  
	engine.GET("/locations", location.ResourceList.GetEndpoint())
	engine.POST("/locations", location.ResourceCreate.GetEndpoint())
	engine.GET("/locations/:id", location.ResourceDetail.GetEndpoint())
	engine.PATCH("/locations/:id", location.ResourceUpdate.GetEndpoint())
	engine.DELETE("/locations/:id", location.ResourceDelete.GetEndpoint())
```

Since we have not defined Pagination and Ordering behaviour, the default ones are used: Limit-Offset pagination and Ordering by the "id" field.

```
GET /locations/
{
  "count": 3,
  "data": [
    {
      "city": "Moscow",
      "country": "Russia",
      "createdAt": "2020-12-23T12:56:11.652588Z",
      "description": "it's cold",
      "id": 2,
      "name": "location2"
    },
    {
      "city": "Texas",
      "country": "USA",
      "createdAt": "2020-12-25T15:37:11.374022Z",
      "description": "TexMex is nice",
      "id": 3,
      "name": "loc3"
    }
  ]
}
```


