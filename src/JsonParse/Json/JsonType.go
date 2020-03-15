package Json

// Author:Boyn
// Date:2020/3/15

const (
	String     = iota // 字符串字面量
	Number            // 数字字面量
	Null              // 空
	StartArray        // '[' 数组开始
	EndArray          // ']' 数组结束
	StartObj          // '{' 对象开始
	EndObj            // '}' 对象结束
	Comma             // , 逗号
	Colon             // : 引号
	Boolean           // 布尔量
	EndDoc            // JSON数据结束
)

func Type2String(t int) string {
	switch t {
	case String:
		return "String"
	case Number:
		return "Number"
	case Null:
		return "NULL"
	case StartArray:
		return "StartArray"
	case EndArray:
		return "EndArray"
	case StartObj:
		return "StartObj"
	case EndObj:
		return "EndObj"
	case Comma:
		return "Comma"
	case Colon:
		return "Colon"
	case Boolean:
		return "Boolean"
	case EndDoc:
		return "EndDoc"
	default:
		return "Unknown"
	}
}
