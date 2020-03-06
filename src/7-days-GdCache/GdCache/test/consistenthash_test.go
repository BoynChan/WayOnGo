package test

import (
	"7-days-GdCache/GdCache/consistenthash"
	"strconv"
	"testing"
)

// Author:Boyn
// Date:2020/3/6

func TestHashing(t *testing.T) {
	hash := consistenthash.New(3, func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	})
	hash.Add("6", "4", "2")
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("k %s v %s", k, v)
		}
	}

	hash.Add("8")
	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("k %s v %s", k, v)
		}
	}
}
