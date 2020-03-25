package main

import (
	"fmt"
	"reflect"
)

// Author:Boyn
// Date:2020/3/23
// in this file, we will use recursive decline to resolve and print the complex type like struct or interface

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T)\n", name, x)
	display(name, reflect.ValueOf(x), 0)
}

// handle complex type
// use recursive decline to resolve Struct,Slice,Map...
// path is to record the name of fields
// v is the real Value in reflect package to represent the value of variable
// depth is to record the depth of recursion and avoid circle dependency problem
func display(path string, v reflect.Value, depth int) {
	// if recursive depth more than 10. it means that the depth is too large so we should just cut if off.
	if depth > 10 {
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s=invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), depth+1)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), depth+1)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key), depth+1)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), depth+1)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), depth+1)
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
