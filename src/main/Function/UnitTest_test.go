package main

import "testing"

/*
单元测试,使用testing.T作为参数
*/
func TestSomething(t *testing.T) {
	area := calArea(5, 6)
	if area != 30 {
		t.Error("面积不等")
	}
}
