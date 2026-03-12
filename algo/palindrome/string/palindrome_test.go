package string

import "testing"

func TestIsPlalindrome(t *testing.T) {
	tests := []struct {
		sep  string
		want bool
	}{
		{"a", true},
		{"hello", false},
		{"abccba", true},
		{"abcdcba", true},
	}
	if IsPlalindrome("") != false {
		t.Errorf(`isPlalindrome(" ")=false`)
	}

	for _, test := range tests {
		if IsPlalindrome(test.sep) != test.want {
			t.Errorf(`isPlalindrome("%s")=false`, test.sep)
		}
	}
}

func TestIsPlalindrome1(t *testing.T) {
	tests := []struct {
		sep  string
		want bool
	}{
		{"a", true},
		{"hello", false},
		{"abccba", true},
		{"abcdcba", true},
	}
	if IsPlalindrome1("") != false {
		t.Errorf(`isPlalindrome1(" ")=false`)
	}

	for _, test := range tests {
		if IsPlalindrome1(test.sep) != test.want {
			t.Errorf(`isPlalindrome1("%s")=false`, test.sep)
		}
	}
}

func TestIsPlalindrome2(t *testing.T) {
	tests := []struct {
		sep  string
		want bool
	}{
		{"a", true},
		{"hello", false},
		{"abccba", true},
		{"abcdcba", true},
	}
	if IsPlalindrome2("") != false {
		t.Errorf(`isPlalindrome1("")=false`)
	}

	for _, test := range tests {
		if IsPlalindrome2(test.sep) != test.want {
			t.Errorf(`isPlalindrome1("%s")=false`, test.sep)
		}
	}
}
