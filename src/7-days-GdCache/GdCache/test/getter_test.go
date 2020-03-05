package test

import (
	. "7-days-GdCache/GdCache"
	"reflect"
	"testing"
)

// Author:Boyn
// Date:2020/3/5

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})
	value, _ := f.Get("123")
	expect := []byte("123")
	if !reflect.DeepEqual(value, expect) {
		t.Fail()
	}
}
