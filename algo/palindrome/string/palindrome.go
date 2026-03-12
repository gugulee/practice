package string

func IsPlalindrome(str string) bool {
	if 0 == len(str) {
		return false
	}

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func IsPlalindrome1(str string) bool {
	if 0 == len(str) {
		return false
	}
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func IsPlalindrome2(str string) bool {
	length := len(str)
	if 0 == len(str) {
		return false
	}
	for i := length/2 - 1; i >= 0; i-- {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
