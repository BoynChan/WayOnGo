package JsonParse

import (
	"JsonParse/Json"
	"fmt"
)

// Author:Boyn
// Date:2020/3/15

type JsonParser struct {
	token    []Json.Token
	parsePtr int
	length   int
}

func NewJsonParser(s string) *JsonParser {
	tokens, _ := parse(s)
	return &JsonParser{
		token:    tokens,
		parsePtr: 0,
		length:   len(tokens),
	}
}

func (j *JsonParser) ParseObject() (map[string]interface{}, error) {
	if j.token[j.parsePtr].TokenType != Json.StartObj {
		return nil, fmt.Errorf("开头错误:不以{开头")
	}
	if j.token[j.parsePtr+1].TokenType == Json.EndObj {
		return make(map[string]interface{}), nil
	}
	result := make(map[string]interface{})
	err := j.parseObject(result)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (j *JsonParser) parseObject(m map[string]interface{}) error {
	var key string
	j.next()
	if j.peek().TokenType != Json.String {
		return fmt.Errorf("解析错误:Key不为String类型")
	}
	key = j.peek().Value.(string)

	j.next()
	if j.peek().TokenType != Json.Colon {
		return fmt.Errorf("解析错误:键值不以\":\"分割")
	}

	j.next()
	// 解析key:后面的value,只有三种可能,一是字面值,二是数组开始的符号 [, 三是对象开始的符号 {
	if isPrimary(j.peek()) {
		// 如果是字符串,数字,布尔值就直接放入map中
		m[key] = j.peek().Value
	} else if j.peek().TokenType == Json.StartArray {
		array := make([]interface{}, 0)
		a, err := j.parseArray(array)
		if err != nil {
			return err
		}
		m[key] = a
	} else if j.peek().TokenType == Json.StartObj {
		result := make(map[string]interface{})
		err := j.parseObject(result)
		if err != nil {
			return err
		}
		m[key] = result
	} else {
		return fmt.Errorf("value解析错误")
	}

	j.next()
	// 解析完value后,解析后面的内容,只有两种可能 一是, 表示后面仍有kv,二是 } 表示结束直接返回
	if j.peek().TokenType == Json.Comma {
		err := j.parseObject(m)
		if err != nil {
			return err
		}
	} else if j.peek().TokenType == Json.EndObj {
		return nil
	} else {
		return fmt.Errorf("结束符错误")
	}
	return nil
}

func (j *JsonParser) parseArray(a []interface{}) ([]interface{}, error) {
	j.next()
	if isPrimary(j.peek()) {
		a = append(a, j.peek().Value)
	} else if j.peek().TokenType == Json.StartArray {
		array := make([]interface{}, 0)
		ar, err := j.parseArray(array)
		if err != nil {
			return nil, err
		}
		a = append(a, ar)
	} else if j.peek().TokenType == Json.StartObj {
		result := make(map[string]interface{})
		err := j.parseObject(result)
		if err != nil {
			return nil, err
		}
		a = append(a, result)
	}
	j.next()

	if j.peek().TokenType == Json.Comma {
		ar, err := j.parseArray(a)
		a = ar
		if err != nil {
			return nil, err
		} else {
			return a, nil
		}
	} else if j.peek().TokenType == Json.EndArray {
		return a, nil
	} else {
		return nil, fmt.Errorf("解析错误")
	}
}

func isPrimary(token Json.Token) bool {
	return token.TokenType == Json.String ||
		token.TokenType == Json.Number ||
		token.TokenType == Json.Boolean ||
		token.TokenType == Json.Null
}

func (j *JsonParser) peek() Json.Token {
	return j.token[j.parsePtr]
}

func (j *JsonParser) next() {
	j.parsePtr++
}

func (j *JsonParser) rollback() {
	j.parsePtr--
}
