package utility

import (
	"fmt"
	"reflect"
)

func ReflectStruct(i interface{})[]interface{}{
	data := make([]interface{} , 0)
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr{
		t = t.Elem()
		v = v.Elem()
	}else{
		Debug("无效的指针")
		return data
	}

	for n :=0; n < t.NumField(); n++{
		fmt.Println(t.Field(n).Name , v.Field(n).Kind())
		data = append(data , t.Field(n))
	}
	return data
}
