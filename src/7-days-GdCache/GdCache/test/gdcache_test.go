package test

import (
	"7-days-GdCache/GdCache"
	"fmt"
	"testing"
)

// Author:Boyn
// Date:2020/3/5

var db = map[string]string{
	"Tom":  "630",
	"Jack": "1630",
	"Sam":  "2630",
	"Bob":  "3630",
	"Carl": "4630",
}

func TestGdGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))
	gd := GdCache.NewGroup("number", 2<<10, GdCache.GetterFunc(func(key string) ([]byte, error) {
		fmt.Printf("DB - [取值] - %s\n", key)
		if v, ok := db[key]; ok {
			if _, ok := loadCounts[key]; !ok {
				loadCounts[key] = 0
			}
			loadCounts[key] += 1
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s 不存在", key)
	}))
	for k, v := range db {
		if view, err := gd.Get(k); err != nil || view.String() != v {
			t.Fail()
		}
		if _, err := gd.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fail()
		}
	}
}
