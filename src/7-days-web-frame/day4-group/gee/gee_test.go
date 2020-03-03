package gee

import (
	"reflect"
	"testing"
)

// Author:Boyn
// Date:2020/3/3

func newTestRouter() *route {
	r := newRoute()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/asserts/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = reflect.DeepEqual(parsePattern("/p/*any"), []string{"p", "*any"})
	ok = reflect.DeepEqual(parsePattern("/p/*name"), []string{"p", "*name"})
	if !ok {
		t.Fatal("解析出错")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/boyn")
	if n == nil {
		t.Fatal("Can not find pattern")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("Wrong Trie parse")
	}
	if ps["name"] != "boyn" {
		t.Fatal("Params injection fail")
	}
}
