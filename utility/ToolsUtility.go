package utility

import (
	"fmt"
	"reflect"
)

func ReflectStruct(i interface{}) []interface{} {
	data := make([]interface{}, 0)
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	} else {
		Debug("无效的指针")
		return data
	}

	for n := 0; n < t.NumField(); n++ {
		fmt.Println(t.Field(n).Name, v.Field(n).Kind())
		data = append(data, t.Field(n))
	}
	return data
}

func GetInt64(n interface{}) int64 {
	var r int64 = 0
	typeName := reflect.TypeOf(n)
	switch typeName.String() {
	case "int":
		r = int64(n.(int))
	case "int32":
		r = int64(n.(int32))
	case "int64":
		r = n.(int64)
	case "byte":
		r = int64(n.(byte))
	}
	return r
}
