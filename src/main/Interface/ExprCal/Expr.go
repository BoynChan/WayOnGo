package ExprCal

import (
	"fmt"
	"math"
)

//表达式父接口
type Expr interface {
	Eval(env Env) float64
}

// Var 表示变量
type Var string

//Var的Eval直接从上下文中寻找结果
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literal 表示数字常量
type literal float64

//literal的Eval直接返回自身
func (l literal) Eval(env Env) float64 {
	return float64(l)
}

// unary 表示一元操作符表达式,op为 +,- x为后面的表达式
type unary struct {
	op rune
	x  Expr
}

// unary的Eval函数先考虑op的值
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	default:
		panic(fmt.Sprintf("Unsupported Operator: %q", u.op))
	}
}

// binary 表示二元操作符,op 为 + - * \ , xy为前后表达式
type binary struct {
	op   rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	default:
		panic(fmt.Sprintf("Unsupported Operator: %q", b.op))
	}
}

// call 表示函数调用表达式 比如 sin(x)
type call struct {
	fn   string // sin cos pow sqrt ...
	args []Expr // 函数中的参数
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	default:
		panic(fmt.Sprintf("Unsupported Operator: %q", c.fn))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "cos":
		return math.Cos(c.args[0].Eval(env))
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
}

// Env 表示将变量映射为数值的上下文
type Env map[Var]float64
