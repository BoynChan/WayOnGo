package JsonParse

import (
	"JsonParse/Json"
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

// Author:Boyn
// Date:2020/3/15
var buf []rune
var length int
var ptr int

func parse(jS string) ([]Json.Token, error) {
	buf = []rune(jS)
	ptr = 0
	length = len(buf)
	tokens := make([]Json.Token, 0)
	for {
		token, err := start()
		if err != nil && err == fmt.Errorf("EOF") {
			return tokens, err
		} else if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
}

func start() (Json.Token, error) {
	c := read()
	for isSpace(c) {
		c = read()
	}
	if isNull(c) {
		return Json.Token{
			Value:     nil,
			TokenType: Json.Null,
		}, nil
	}
	if c == ',' {
		return Json.Token{
			Value:     ",",
			TokenType: Json.Comma,
		}, nil
	}
	if c == ':' {
		return Json.Token{
			Value:     ":",
			TokenType: Json.Colon,
		}, nil
	}
	if c == '{' {
		return Json.Token{
			Value:     "{",
			TokenType: Json.StartObj,
		}, nil
	}
	if c == '}' {
		return Json.Token{
			Value:     "}",
			TokenType: Json.EndObj,
		}, nil
	}
	if c == '[' {
		return Json.Token{
			Value:     "[",
			TokenType: Json.StartArray,
		}, nil
	}
	if c == ']' {
		return Json.Token{
			Value:     "]",
			TokenType: Json.EndArray,
		}, nil
	}
	if isTrue(c) {
		return Json.Token{
			Value:     true,
			TokenType: Json.Boolean,
		}, nil
	}
	if isFalse(c) {
		return Json.Token{
			Value:     false,
			TokenType: Json.Boolean,
		}, nil
	}
	if c == '"' {
		return readString()
	}
	if isNum(c) {
		rollback(1)
		n, err := readNum()
		rollback(1)
		return n, err
	}
	if c == -1 {
		return Json.Token{}, fmt.Errorf("EOF")
	}
	return Json.Token{}, fmt.Errorf("JSON解析错误")

}

func readNum() (Json.Token, error) {
	buffer := bytes.Buffer{}
	c := read()
	if c == '-' {
		buffer.WriteRune(c)
		c = read()
		if c == '0' {
			buffer.WriteRune(c)
			numAppend(&buffer)
		} else if unicode.IsDigit(c) {
			for unicode.IsDigit(c) {
				buffer.WriteRune(c)
				c = read()
			}
			rollback(1)
			numAppend(&buffer)
		} else {
			return Json.Token{}, fmt.Errorf("number解析错误")
		}
	} else if c == '0' {
		buffer.WriteRune(c)
		numAppend(&buffer)
	} else if unicode.IsDigit(c) {
		for unicode.IsDigit(c) {
			buffer.WriteRune(c)
			c = read()
		}
		rollback(1)
		numAppend(&buffer)
	}
	num, _ := strconv.Atoi(buffer.String())
	return Json.Token{
		Value:     num,
		TokenType: Json.Number,
	}, nil
}

func numAppend(buffer *bytes.Buffer) {
	c := read()
	if c == '.' {
		buffer.WriteRune(c)
		c = read()
		for unicode.IsDigit(c) {
			buffer.WriteRune(c)
			c = read()
		}
		rollback(1)
		expAppend(buffer)
	} else {
		expAppend(buffer)
	}
}

func expAppend(buffer *bytes.Buffer) {
	c := read()
	if c == 'e' || c == 'E' {
		buffer.WriteRune(c)
		c = read()
		for unicode.IsDigit(c) {
			buffer.WriteRune(c)
			c = read()
		}
	}
	rollback(1)
}

func readString() (Json.Token, error) {
	buffer := bytes.Buffer{}
	for {
		c := read()
		if b, err := isEscape(c); err == nil && b {
			if c == 'u' {
				buffer.WriteRune('\\')
				buffer.WriteRune(c)
				for i := 0; i < 4; i++ {
					c = read()
					if isHex(c) {
						buffer.WriteRune(c)
					} else {
						return Json.Token{}, fmt.Errorf("\\u解析错误")
					}
				}
			} else {
				rollback(1)
				buffer.WriteRune(read())
			}
		} else if err != nil {
			return Json.Token{}, err
		} else if c == '"' {
			return Json.Token{
				Value:     buffer.String(),
				TokenType: Json.String,
			}, nil
		} else if c == '\r' || c == '\n' {
			return Json.Token{}, fmt.Errorf("非法输入")
		} else {
			buffer.WriteRune(c)
		}
	}
}

func isHex(c rune) bool {
	return (c > '0' && c < '9') || (c > 'A' && c < 'E')
}

func isEscape(c rune) (bool, error) {
	if c == '\\' {
		c = read()
		if c == '"' || c == '\\' || c == '/' || c == 'b' || c == 'f' || c == 'n' || c == 't' || c == 'r' || c == 'u' {
			return true, nil
		} else {
			return false, fmt.Errorf("escape解析错误")
		}

	} else {
		return false, nil
	}
}

func read() rune {
	if ptr >= length {
		return -1
	}
	b := buf[ptr]
	ptr += 1
	return b
}

func rollback(i int) {
	ptr -= i
}

func isTrue(c rune) bool {
	isCapital := false
	if c == 't' {
		isCapital = false
		c = read()
	} else if c == 'T' {
		isCapital = true
		c = read()
	} else {
		return false
	}
	if (!isCapital && c == 'r') || (isCapital && c == 'R') {
		c = read()
	} else {
		rollback(1)
		return false
	}

	if (!isCapital && c == 'u') || (isCapital && c == 'U') {
		c = read()
	} else {
		rollback(2)
		return false
	}

	if (!isCapital && c == 'e') || (isCapital && c == 'E') {
		return true
	} else {
		rollback(3)
		return false
	}
}

func isFalse(c rune) bool {
	isCapital := false
	if c == 'f' {
		isCapital = false
		c = read()
	} else if c == 'F' {
		isCapital = true
		c = read()
	} else {
		return false
	}
	if (!isCapital && c == 'a') || (isCapital && c == 'A') {
		c = read()
	} else {
		rollback(1)
		return false
	}

	if (!isCapital && c == 'l') || (isCapital && c == 'L') {
		c = read()
	} else {
		rollback(2)
		return false
	}

	if (!isCapital && c == 's') || (isCapital && c == 'S') {
		c = read()
	} else {
		rollback(3)
		return false
	}

	if (!isCapital && c == 'e') || (isCapital && c == 'E') {
		return true
	} else {
		rollback(4)
		return false
	}
}

func isNull(c rune) bool {
	isCapital := false
	if c == 'n' {
		isCapital = false
		c = read()
	} else if c == 'N' {
		isCapital = true
		c = read()
	} else {
		return false
	}
	if (!isCapital && c == 'u') || (isCapital && c == 'U') {
		c = read()
	} else {
		rollback(1)
		return false
	}

	if (!isCapital && c == 'u') || (isCapital && c == 'U') {
		c = read()
	} else {
		rollback(2)
		return false
	}

	if (!isCapital && c == 'u') || (isCapital && c == 'U') {
		return true
	} else {
		rollback(3)
		return false
	}
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\n' || c == '\t'
}

func isNum(c rune) bool {
	return unicode.IsDigit(c) || c == '-'
}
