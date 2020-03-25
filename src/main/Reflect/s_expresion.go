package main

import (
	"bytes"
	"fmt"
	"reflect"
)

// Author:Boyn
// Date:2020/3/24
// we will use reflect package to encode any go variable into s expresion format
// s expresion is an old expresion language which use the lisp-like grammar

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := encode(&buf, reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// the function can encode any go object (but not channel or function) into s expresion
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Bool:
		if v.Bool() == true {
			fmt.Fprintf(buf, "t")
		} else {
			fmt.Fprintf(buf, " ")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, _ = fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_, _ = fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		_, _ = fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			// encode the key
			fmt.Fprintf(buf, "( %s ", v.Type().Field(i).Name)
			// encode the value
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			// encode the key
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			// encode the value
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	default:
		return fmt.Errorf("unsupported type:%s", v.Type())
	}
	return nil
}
