package sstruct

import (
	"reflect"
)

// ToFieldValueMap struct to map
func ToFieldValueMap(data interface{}) map[string]interface{} {
    list := map[string]interface{}{}
    if data == nil {
        return list
    }
    v := reflect.ValueOf(data)
    typeV := v.Type()
    for i := 0; i < v.NumField(); i++ {
        list[typeV.Field(i).Name] = v.Field(i).Interface()
    }
    
    return list
}


// ToMapByTag  struct to map using tag as key
func ToMapByTag(data interface{}, tag string) map[string]interface{} {
    list := map[string]interface{}{}
    if data == nil {
        return list
    }
    v := reflect.ValueOf(data)
    typeV := v.Type()
    
    for i := 0; i < v.NumField(); i++ {
        fieldTag, ok := typeV.Field(i).Tag.Lookup(tag)
        if !ok {
            fieldTag = typeV.Field(i).Name
        }
        list[fieldTag] = v.Field(i).Interface()
    }
    return list
}

// ToMapByTagJson  struct to map using json tag as key
func ToMapByTagJson(data interface{}) map[string]interface{} {
    return ToMapByTag(data, "json")
}
