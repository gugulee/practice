package string

import "testing"

func TestIsPlalindrome1(t *testing.T) {
	strs := []string{"hello", "abccba", "abcdcba"}
	for _, str := range strs {
		if !IsPlalindrome1(str) {
			t.Errorf(`isPlalindrome1("%s")=false`, str)
		}else {
			t.Logf(`isPlalindrome1("%s")=true`, str)
		}
	}
}


func TestIsPlalindrome2(t *testing.T) {
	strs := []string{"hello", "abccba", "abcdcba"}
	for _, str := range strs {
		if !IsPlalindrome2(str) {
			t.Errorf(`isPlalindrome1("%s")=false`, str)
		}else {
			t.Logf(`isPlalindrome1("%s")=true`, str)
		}
	}
}