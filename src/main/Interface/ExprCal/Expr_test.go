package ExprCal

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A/pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x,3)+pow(y,3)", Env{"x": 2, "y": 3}, "35"},
		{"pow(x,3)+pow(y,3)", Env{"x": 1, "y": 2}, "9"},
		{"5/9*(F-32)", Env{"F": -40}, "-40"},
		{"5/9*(F-32)", Env{"F": 32}, "0"},
		{"5/9*(F-32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		//通过Parse函数解析表达式
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q,want %q\n", test.expr, test.env, got, test.want)
		}
	}
}
