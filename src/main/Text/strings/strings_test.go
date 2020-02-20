package strings

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

// Author:Boyn
// Date:2020/2/20

// 记录字符串的操作

func TestCompare(t *testing.T) {
	a := "gopher"
	b := "hello world"
	c := "HELLO WORLD"

	// Compare函数 当a=b 时 返回0,当a<b时, 返回-1,a>b时返回+1,大于小于的关系是按照字典序逐个比对实现的
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(c, a))

	// EqualFold返回一个bool,用于比较两个字符串在忽略大小写后是否相等
	fmt.Println(strings.EqualFold(a, b))
	fmt.Println(strings.EqualFold(c, b))
}

func TestContain(t *testing.T) {
	// Contains s的子串中是否存在substr
	fmt.Println(strings.Contains("abcdefg", "abc"))

	// ContainsAny chars中是否有任一字符存在于s中
	fmt.Println(strings.ContainsAny("abcdefg", "werty"))

	// ContainsRune r代码点是否存在于s中
	fmt.Println(strings.ContainsRune("我是谁", '我'))
}

func TestCount(t *testing.T) {
	// Count查看子串无重叠出现次数(如fivevev,vev只算一次)
	fmt.Println(strings.Count("I am a student.", "a"))

}

func TestSplit(t *testing.T) {
	// Fields和FieldsFunc
	// Fields用于分割字符串,分割点是常见的间隔符,如空格,\t,\n
	fmt.Printf("%q\n", strings.Fields("   foo   bar   baz"))

	// FieldsFunc可以传入一个函数用于规定分割的字符串
	fmt.Printf("%q\n", strings.FieldsFunc("   foo   bar   baz", func(r rune) bool {
		return r == ' '
	}))
	// Split和SplitAfter都会将字符串以规定的字符进行分割,After会保留分隔符
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
}

func TestPrefixSuffix(t *testing.T) {
	// 是否以某个前缀开头
	fmt.Println(strings.HasPrefix("abc", "a"))
	fmt.Println(strings.HasPrefix("abc", "b"))

	// 是否以某个后缀结尾
	fmt.Println(strings.HasSuffix("abc", "c"))
	fmt.Println(strings.HasSuffix("abc", "b"))
}

func TestIndex(t *testing.T) {
	// Index系列有Index,IndexAny,IndexFunc和IndexRune,其使用方法和返回值大同小异,在这里用IndexFunc举例
	// 同样地,还有LastIndex系列,也是一样的方法.就不举例了
	fmt.Println(strings.IndexFunc("Hello,世界", func(r rune) bool {
		// 返回是否汉字
		return unicode.Is(unicode.Han, r)
	}))
}

func TestJoin(t *testing.T) {
	fmt.Println(strings.Join([]string{"a", "b"}, ","))
	// Join函数可以将字符串连接起来
}

func TestMapReplace(t *testing.T) {
	// Map函数会将s中的所有字符按照mapping的规则做替换
	// 这个mapping函数将大写转换为小写
	mapping := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r + 32
		default:
			return r
		}
	}
	fmt.Println(strings.Map(mapping, "abcdEFG"))
}

func TestReplace(t *testing.T) {
	// Replace参数为 源字符串,旧字串,新字串,替换个数(从左到右),当替换个数<0时,即替换全部
	fmt.Println(strings.Replace("i am boyn. i am a student. i wrote this.", "i", "I", 1))
	fmt.Println(strings.Replace("i am boyn. i am a student. i wrote this.", "i", "I", 2))
	fmt.Println(strings.Replace("i am boyn. i am a student. i wrote this.", "i", "I", -1))
	fmt.Println(strings.ReplaceAll("i am boyn. i am a student. i wrote this.", "i", "I"))
}

func TestLowerUpper(t *testing.T) {
	fmt.Println(strings.ToUpper("abcd"))
	fmt.Println(strings.ToLower("ABCD"))
}

func TestTitle(t *testing.T) {
	// 将每个单词的首字母大写
	fmt.Println(strings.Title("hello world"))
}
