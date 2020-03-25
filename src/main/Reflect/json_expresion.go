package main

import (
	"bytes"
	"fmt"
	"reflect"
)

//Author: Boyn
//Date: 2020/3/25
// this file contains jsonEncode function and JsonMarshal function
// which will jsonEncode any go object into json string

// this function is a wrapper
// it accept a interface{} v represent arbitrary go variable
func JsonMarshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := jsonEncode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// this function is to encode v into json string and store in buf Buffer
func jsonEncode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")
	case reflect.Bool:
		if v.Bool() == true {
			fmt.Fprintf(buf, "true")
		} else {
			fmt.Fprintf(buf, "false")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, _ = fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_, _ = fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		_, _ = fmt.Fprintf(buf, "%.3f", v.Float())
	case reflect.String:
		_, _ = fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return jsonEncode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := jsonEncode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')
	case reflect.Struct:
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			} // jsonEncode the key
			fmt.Fprintf(buf, "\"%s\":", v.Type().Field(i).Name)
			// jsonEncode the value
			if err := jsonEncode(buf, v.Field(i)); err != nil {
				return err
			}
			//buf.WriteByte('')
		}
		buf.WriteByte('}')
	case reflect.Map:
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(',')
			}
			//buf.WriteByte('(')
			// jsonEncode the key
			if err := jsonEncode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			// jsonEncode the value
			if err := jsonEncode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			//buf.WriteByte(')')
		}
		buf.WriteByte('}')
	default:
		return fmt.Errorf("unsupported type:%s", v.Type())
	}
	return nil
}
