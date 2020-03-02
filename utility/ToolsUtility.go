package utility

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
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

func GetRandNum(num ...int64) int64 {
	var min, max int64 = 0, 0
	if len(num) == 1 {
		max = num[0]
	} else {
		min = num[0]
		max = num[1]
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}
