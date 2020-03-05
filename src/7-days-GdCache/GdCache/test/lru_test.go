package test

import (
	"7-days-GdCache/GdCache/lru"
	"testing"
)

// Author:Boyn
// Date:2020/3/5
type Integer int

func (Integer) Len() int {
	return 4
}

func NewCache(cap int64) *lru.Cache {
	return lru.New(cap, nil)
}

func TestAdd(t *testing.T) {
	// 内存足够
	cache := NewCache(15)
	cache.Add("1", Integer(1))
	cache.Add("2", Integer(2))
	cache.Add("3", Integer(3))
	if !cache.IsExist("1") {
		t.Fail()
	}
	if !cache.IsExist("2") {
		t.Fail()
	}
	if !cache.IsExist("3") {
		t.Fail()
	}

	// 内存不够,会淘汰1
	cache = NewCache(10)
	cache.Add("1", Integer(1))
	cache.Add("2", Integer(3))
	cache.Add("3", Integer(2))
	if cache.IsExist("1") {
		t.Fail()
	}
	if !cache.IsExist("2") {
		t.Fail()
	}
	if !cache.IsExist("3") {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	cache := NewCache(15)
	cache.Add("1", Integer(1))
	cache.Add("2", Integer(2))
	cache.Add("3", Integer(3))
	if ele, ok := cache.Get("1"); ele != Integer(1) || !ok {
		t.Fail()
	}
	if ele, ok := cache.Get("2"); ele != Integer(2) || !ok {
		t.Fail()
	}
	if ele, ok := cache.Get("3"); ele != Integer(3) || !ok {
		t.Fail()
	}
}
