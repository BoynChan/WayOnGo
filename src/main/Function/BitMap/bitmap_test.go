package BitMap

import (
	"fmt"
	"testing"
)

func TestIntSet(t *testing.T) {
	x, y := new(IntSet), new(IntSet)
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(5)
	if x.Has(5) == false {
		t.Error("Has函数出错")
	}
	if x.String() != "{1 2 3 5}" {
		t.Error("String()函数出错")
	}

	y.Add(9)
	y.Add(35)
	y.Add(13)
	y.Add(164)

	if y.Has(164) == false {
		t.Error("动态扩容出错")
	}

	x.UnionWith(y)

	if x.String() != "{1 2 3 5 9 13 35 164}" {
		t.Error("并集函数出错")
	}
}

func TestBitMapCopy(t *testing.T) {
	s := new(IntSet)
	s.Add(1)
	s.Add(65)
	y := s.Copy()
	if &s.words[0] == &y.words[0] {
		t.Error("深拷贝函数出错,引用到了同一个地址")
	}
	for i, v := range s.words {
		if v != y.words[i] {
			t.Error("值出错,深拷贝后值不一致")
		}
	}
}

func TestBitMapLen(t *testing.T) {
	s := new(IntSet)
	s.Add(1)
	if s.Len() != 1 {
		t.Fail()
	}
	s.Add(2)
	if s.Len() != 2 {
		t.Fail()
	}
}

func TestBitMapElem(t *testing.T) {
	s := new(IntSet)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	for _, v := range s.Elems() {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
