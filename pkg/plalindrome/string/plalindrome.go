package string

func IsPlalindrome1(str string) bool {
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
	for i := length/2-1; i >=0; i-- {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
