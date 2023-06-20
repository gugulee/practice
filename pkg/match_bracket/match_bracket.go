package match_bracket

var bracketRelation = map[string]string{"(": ")", "[": "]", "{": "}"}
var rightBracket = map[string]struct{}{
	")": {},
	"]": {},
	"}": {},
}
var leftBracket = map[string]struct{}{
	"(": {},
	"[": {},
	"{": {},
}

func MatchBracket1(s string) bool {
    n := len(s)
    if n % 2 == 1 {
        return false
    }
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    stack := []byte{}
    for i := 0; i < n; i++ {
        if pairs[s[i]] > 0 {
            if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}

func MatchBracket(expression string) bool {
	brackets := make([]string, 0)
	for _, e := range expression {
		s := string(e)
		if !isBracket(s) {
			continue
		}

		// push e into stack if e is a left bracket
		if isLeftBracket(s) {
			brackets = append(brackets, s)
		} else if isRightBracket(s) {
			if len(brackets) == 0 {
				return false
			}

			// pop the top of stack if e is a right bracket
			top := brackets[len(brackets)-1]
			brackets = brackets[:len(brackets)-1]

			if v := bracketRelation[top]; v != s {
				return false
			}
		}
	}

	return len(brackets) == 0
}

func isBracket(a string) bool {
	return isLeftBracket(a) || isRightBracket(a)
}

func isLeftBracket(a string) bool {
	_, ok := leftBracket[a]
	return ok
}

func isRightBracket(a string) bool {
	_, ok := rightBracket[a]
	return ok
}
