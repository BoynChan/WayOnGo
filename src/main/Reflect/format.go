package main

import (
	"reflect"
	"strconv"
)

// Author:Boyn
// Date:2020/3/23
// 使用reflect来实现将任意值转换成字符串

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// in this function, we can print the value of atom type like int,uint... but not struct or interface type
// these type is called complex type which we will handle it later
// and for the reference type, it's not friendly enough either
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'e', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		return strconv.FormatFloat(real(v.Complex()), 'e', -1, 64) + "i" + strconv.FormatFloat(imag(v.Complex()), 'e', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}
