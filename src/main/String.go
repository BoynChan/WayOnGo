package main

import "fmt"

func main() {
	var str1 = "博言技术"
	str2 := "https://boyn.top"
	//可以使用+=拼接字符串
	str1 += str2
	fmt.Println(str1)

	longStr := `
这是一段长文本
其中有多行
这种称为字符串字面量`
	fmt.Println(longStr)
	// 格式化说明符%c用于表示字符，当和字符配合使用时，
	// %v或%d会输出用于表示该字符的整数，%U输出格式为 U+hhhh 的字符串。
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00001234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}
