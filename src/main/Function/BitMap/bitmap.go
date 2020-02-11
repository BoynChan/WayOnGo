package BitMap

import (
	"bytes"
	"fmt"
)

//以slice作为一个位图的结构
//其中,将位图以64为单位进行分组,一个uint64有64位,可以表示64个布尔量
//所以我们这个位图以uint64[]进行底层结构
type IntSet struct {
	words []uint64
}

// 要检查某个位置的变量是否存在,我们可以通过位运算来计算
// 首先,我们要算出入参所属的组数与组中的个数
// 之后要检查入参对应的组数是否小于位图的组数
// 然后检查对应位置的位是否为0,具体为按组取出uint64后,使用位运算,向左移动到对应的位数之后
// 进行检查该位是否为1
func (s *IntSet) Has(x int) bool {
	group, index := x/64, uint(x%64)
	return group < len(s.words) && s.words[group]&(1<<index) != 0
}

// 对于添加一个元素到位图中的操作,我们同样需要取组数和组中个数
// 当组数小于位图组数的时候,我们就要进行扩容,直到组数=位图组数+1为止
// 然后使用或运算,将组中个数那一位设为1
func (s *IntSet) Add(x int) {
	group, index := x/64, uint(x%64)
	for group >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[group] |= 1 << index
}

// 跟另一个位图进行并集运算
// 需要考虑的场景就是长度问题了
// 我们可以对t进行遍历
// 当t长度大于s,就让s动态扩容,并且扩容的内容是t数组后面的内容
// 当s长度大于t,就可以不需要对t的遍历过程做特别处理了
// 他们重合的那一部分直接用或运算符进行赋值
func (s *IntSet) UnionWith(t *IntSet) {
	sLen := len(s.words)
	for index, word := range t.words {
		if index < sLen {
			s.words[index] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// 实现与另外一个位图进行交集运算
// 不需要考虑长度问题,因为是以s集合为主的
// 所以只用考虑s长度内的问题即可
// 我们只用把那边的或运算换成与运算即可
func (s *IntSet) IntersectWith(t *IntSet) {
	for index, word := range t.words {
		s.words[index] &= word
	}
}

// 实现与另一个位图进行差集运算
// 元素出现在s中但没有出现在t中
// 思路是先保存s集合原有的元素
// 然后将s集合与t进行异或,再将结果与原来进行与运算
// e.g.  s : 001101  t: 001011   s^t = 000110  s'&s=000100
func (s *IntSet) DifferenceWith(t *IntSet) {
	for index, word := range t.words {
		temp := s.words[index]
		temp ^= word
		s.words[index] &= temp
	}
}

// 实现并差集运算
// 元素出现在s中但没有出现在t中,或者元素出现在t中但没有出现在s中
// 思路是对这两个集合互相求差集,然后求并集
func (s *IntSet) SymmetricDifference(t *IntSet) {
	temp := t.Copy()
	temp.DifferenceWith(s)
	s.DifferenceWith(t)
	s.UnionWith(temp)
}

// String()方法以数字的方式打印出有哪一位是被标记为1的
// 如'{1 2 3}'
func (s *IntSet) String() string {
	buf := new(bytes.Buffer)
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				//写入了第一个元素后再开始加空格
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 返回位图的个数,即遍历位图,寻找1的个数
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				count++
			}
		}
	}
	return count
}

// 删除某一位的元素
// 其方法,首先要获取某一位元素的取反值
// 然后跟对应组的元素进行与运算
func (s *IntSet) Remove(x int) {
	group, index := x/64, uint(x%64)
	if group > len(s.words) {
		return
	}
	//在这里为了避免溢出错误,我们要先将1<<index转为uint64类型再按位取反
	s.words[group] &= ^uint64(1 << index)
}

// 清空所有元素
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

// 返回一个位图的深拷贝
func (s *IntSet) Copy() *IntSet {
	t := new(IntSet)
	// 在append函数中,由于t.words长度为0,所以添加s.words的时候会动态开辟空间
	// 所以符合深拷贝对于不同地址的要求
	t.words = append(t.words, s.words...)
	return t
}

// 使用变参参数来添加一组数字到位图中
// 遍历切片,一个个进行添加即可
func (s *IntSet) AddAll(a ...int) {
	for _, v := range a {
		s.Add(v)
	}
}

// 以切片的形式返回集合中的所有元素
func (s *IntSet) Elems() []int {
	result := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				result = append(result, 64*i+j)
			}
		}
	}
	return result
}
