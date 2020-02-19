package fmt

import "fmt"

// Author:Boyn
// Date:2020/2/19
type user struct {
	name string
}

func fmtPrintf() {
	u := user{name: "Jack"}

	fmt.Printf("%+v\n", u)       // 格式化输出结构
	fmt.Printf("%#v\n", u)       // 输出值的Go表示方法
	fmt.Printf("%T\n", u)        // 输出值的类型的Go表示方法
	fmt.Printf("%t\n", true)     // 输出值的true或者false
	fmt.Printf("%b\n", 1024)     // 输出二进制表示
	fmt.Printf("%c\n", 11111111) // 数值对应的Unicode编码
	fmt.Printf("%d\n", 10)       // 十进制表示
	fmt.Printf("%o\n", 8)        // 八进制表示
	fmt.Printf("%q\n", 22)       // 转为16进制并加上单引号
	fmt.Printf("%x\n", 1223)     // 16进制,字母用a-f表示
	fmt.Printf("%X\n", 1223)     // 16进制,字母用A-F表示
	fmt.Printf("%U\n", 1233)     // Unicode表示
	fmt.Printf("%b\n", 12.34)    // 无小数部分,两位指数的科学计数法
	fmt.Printf("%e\n", 12.345)   // 科学计数法表示,e表示
	fmt.Printf("%E\n", 12.345)   // 科学计数法表示,E表示
	fmt.Printf("%f\n", 12.345)   // 有小数部分,无指数
	fmt.Printf("%g\n", 12.345)   // 根据实际情况使用%e或者%f输出
	fmt.Printf("%G\n", 12.345)   // 根据实际情况使用%E或者%f输出

	fmt.Printf("%s\n", "qwetyu")  // 直接输出字符串
	fmt.Printf("%q\n", "qwertyu") // 双括号括起来的字符串
	fmt.Printf("%x\n", "qwertyu") // 每个字节用16进制表示,a-f
	fmt.Printf("%X\n", "qwertyu") // 每个字节用16进制表示,A-F
}
