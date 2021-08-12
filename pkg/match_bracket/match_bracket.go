package match_bracket

import (
	set "github.com/deckarep/golang-set"
	arraystack "github.com/practice/pkg/stack/array"
	"strings"
	"text/scanner"
)

var rightBracket = set.NewSet(")", "]", "}")
var leftBracket = set.NewSet("(", "[", "{")
var bracketRelation = map[string]string{"(": ")", "[": "]", "{": "}"}

func MatchBracket(expression string) bool {
	Brackets := arraystack.New()
	sc := new(scanner.Scanner)
	sc.Init(strings.NewReader(expression))
	sc.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	var tok rune
	for tok != scanner.EOF {
		tok = sc.Scan()
		e := sc.TokenText()

		// skip empty string
		if 0 == len(e) {
			continue
		}

		// push e into stack if e is a left bracket,
		// check e and the top of stack if e is a right bracket:
		// 1. if match(bracketRelation[stack.top] = e), scan the remainder of the string
		// 2. not match, return false
		if leftBracket.Contains(e) {
			Brackets.Push(e)
		} else if rightBracket.Contains(e) {
			// no left bracket, return false
			if Brackets.IsEmpty() {
				return false
			}

			top := Brackets.Pop().(string)
			if v := bracketRelation[top]; e != v {
				return false
			}
		}

	}

	// if stack is empty, return true
	return Brackets.IsEmpty()
}
