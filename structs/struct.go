package structs

import (
	"reflect"
)

// ToFieldValueMap struct to map
// Deprecated: please using sstruct.ToFieldValueMap
func ToFieldValueMap(i interface{}) map[string]interface{} {
	v := reflect.ValueOf(i)
	typeV := v.Type()
	list := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		list[typeV.Field(i).Name] = v.Field(i).Interface()
	}
	return list
}
