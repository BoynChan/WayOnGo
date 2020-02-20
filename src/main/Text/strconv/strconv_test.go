package strconv

import (
	"fmt"
	"strconv"
	"testing"
)

// Author:Boyn
// Date:2020/2/20

// strconv包主要是字符串与基本数据类型之间的转换
// strconv有两个主要的错误,ErrRange表示超过值能表示的最大范围,ErrSyntax表示语法错误

func TestAtoI(t *testing.T) {
	// base表示进制,bitSize表示整数取值范围
	num, _ := strconv.ParseInt("-12345", 10, 16)
	fmt.Println(num)

	unum, _ := strconv.ParseUint("12345", 10, 16)
	fmt.Println(unum)

	// 简化版的ParseInt
	atoi, _ := strconv.Atoi("-12345")
	fmt.Println(atoi)
}

func TestFormat(t *testing.T) {
	fmt.Println(strconv.FormatInt(12345, 10))
	fmt.Println(strconv.FormatUint(uint64(12345), 10))
}

func TestStringBool(t *testing.T) {
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
	fmt.Println(strconv.FormatBool(true))
}

func TestFloat(t *testing.T) {
	float, _ := strconv.ParseFloat("12345.123", 32)
	fmt.Println(float)
	// fmt为表示方法,e E f有效数字用小数点之后的位数,g G表示所有有效数字
	// prec为有效数字
	strconv.FormatFloat(12345.1234, 'g', 3, 32)

}
