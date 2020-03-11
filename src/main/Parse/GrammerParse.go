package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Author:Boyn
// Date:2020/3/11
// 在这部分中，你将使用图转移算法手工实现一个小型的词法分析器。
//
//* 分析器的输入：存储在文本文件中的字符序列，字符取自ASCII字符集。文件中可能包括四种记号：关键字if、符合C语言标准的标识符、空格符、回车符\n。
//
//* 分析器的输出：打印出所识别的标识符的种类、及行号、列号信息。
//
//【示例】对于下面的文本文件：
//
//ifx if iif       if
//
//iff     if
//
//你的输出应该是：
//
//ID(ifx) (1, 1)
//
//IF        (1, 4)
//
//ID(iif)  (1, 8)
//
//IF       (1, 13)
//
//ID(iff) (2, 1)
//
//IF       (2, 8)

func main() {
	file, err := os.Open("F:\\Code\\Go\\LearningGo\\src\\main\\Parse\\text")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ifSm := New([]int{0, 1, 2, 3}, all)

	output := make([]string, 0)
	temp := make([]byte, 0)
	var row, col int
	for b, err := ifSm.next(); err == nil; b, err = ifSm.next() {
		switch ifSm.state {
		case 0:
			if isI(b) {
				row = ifSm.row
				col = ifSm.col
				ifSm.t(2)
				temp = append(temp, b)
				break
			}
			if isWord(b) {
				row = ifSm.row
				col = ifSm.col
				ifSm.t(1)
				temp = append(temp, b)
				break
			}
		case 1:
			if isSpace(b) {
				s := fmt.Sprintf("ID(%s) (%d,%d)", string(temp), row, col)
				temp = make([]byte, 0)
				output = append(output, s)
				ifSm.t(0)
				break
			}
			if isWord(b) {
				temp = append(temp, b)
				break
			}
		case 2:
			if isF(b) {
				temp = append(temp, b)
				ifSm.t(3)
				break
			}
			if isWord(b) {
				temp = append(temp, b)
				ifSm.t(1)
				break
			}
		case 3:
			if isSpace(b) {
				temp = append(temp, b)
				s := fmt.Sprintf("IF (%d,%d)", row, col)
				temp = make([]byte, 0)
				output = append(output, s)
				ifSm.t(0)
				break
			}
			if isWord(b) {
				temp = append(temp, b)
				ifSm.t(1)
				break
			}
		}
	}
	if ifSm.state == 1 {
		s := fmt.Sprintf("ID(%s) (%d,%d)", string(temp), row, col)
		output = append(output, s)
	}
	if ifSm.state == 3 {
		s := fmt.Sprintf("IF (%d,%d)", row, col)
		output = append(output, s)
	}

	for _, v := range output {
		fmt.Println(v)
	}
}

type State interface {
	t(stateTransferTo int) // 转移函数T : Q×Σ → P(Q)
	next() (byte, error)
	rollBack()
}

type NFA struct {
	readLength int    // 已经读取的长度
	length     int    // 文本长度
	row, col   int    // 读取的行号,列号
	Q          []int  // 状态的有限集合Q
	state      int    // 现在的状态
	Sigma      []byte // 输入符号的有限集合Σ
}

type IfSM struct {
	NFA
}

func New(Q []int, Sigma []byte) *IfSM {
	return &IfSM{NFA{
		readLength: 0,
		length:     len(Sigma),
		row:        1,
		col:        0,
		Q:          Q,
		Sigma:      Sigma,
	}}
}

func (i *IfSM) t(stateTransferTo int) {
	i.state = stateTransferTo
}

func (i *IfSM) next() (byte, error) {
	if i.readLength == i.length {
		return 0, fmt.Errorf("读取结束")
	}
	b := i.Sigma[i.readLength]
	if b == '\n' {
		i.readLength += 1
		i.row++
		i.col = 0
		return ' ', nil
	}

	b = i.Sigma[i.readLength]
	i.readLength += 1
	i.col++
	return b, nil
}

func (i *IfSM) rollBack() {
	panic("implement me")
}

func isI(b byte) bool {
	return b == 'i'
}

func isF(b byte) bool {
	return b == 'f'
}

func isWord(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isSlashN(b byte) bool {
	return b == '\n'
}

func isSpace(b byte) bool {
	return b == ' '
}
