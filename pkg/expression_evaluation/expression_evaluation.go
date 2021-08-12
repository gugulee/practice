package expression_evaluation

import (
	"strconv"
	"strings"
	"text/scanner"

	arraystack "github.com/practice/pkg/stack/array"
)

// the priority of operator is "+" = "-" < "*" = "/"
var operatorPriority = map[string]int{"nil": 0, "+": 1, "-": 1, "*": 2, "/": 2}

// operatorPriority compare the priority of operator,
// return true if o2 > o1, else return false
func OperatorPriority(o1, o2 string) bool {
	return operatorPriority[o1] < operatorPriority[o2]
}

// ExpressionEvaluation calculate value of the expression like 3+5*8-6
func ExpressionEvaluation(expression string) (int, error) {
	numeric := arraystack.New()
	operator := arraystack.New()
	sc := new(scanner.Scanner)
	sc.Init(strings.NewReader(expression))
	sc.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	var tok rune
	for tok != scanner.EOF {
		tok = sc.Scan()
		e := sc.TokenText()

		if 0 == len(e) {
			continue
		}

		// e is a operator,
		// push e into operator stack if the priority of e is bigger than the top of operator stack,
		// else calculate with two element from numeric stack and one element from operator stack
		if _, ok := operatorPriority[e]; ok {
			for {
				if operator.IsEmpty() {
					operator.Push(e)
					break
				}

				top := operator.Top().(string)

				if OperatorPriority(top, e) {
					operator.Push(e)
					break
				} else {
					numeric.Push(calculate(numeric.Pop().(int), numeric.Pop().(int), operator.Pop().(string)))
				}
			}
		} else { // e is a numeric, push e into stack directly
			i, err := strconv.Atoi(e)
			if nil != err {
				return 0, err
			}
			numeric.Push(i)
		}
	}

	for ! operator.IsEmpty() {
		numeric.Push(calculate(numeric.Pop().(int), numeric.Pop().(int), operator.Pop().(string)))
	}

	return numeric.Top().(int), nil
}

// calculate calculate the value, "second" operator(+,-,*,/) "first"
func calculate(first, second int, operator string) int {
	switch operator {
	case "+":
		return second + first
	case "-":
		return second - first
	case "*":
		return second * first
	case "/":
		return second / first
	default:
		return 0
	}

}
