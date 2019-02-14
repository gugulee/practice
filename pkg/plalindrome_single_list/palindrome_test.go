package plalindrome_single_list

import (
	"testing"

	single "github.com/practice/pkg/slngle_list"
)

func TestIsPlalindrome1(t *testing.T) {
	tests := []struct {
		sep  string
		want bool
	}{
		{"", false},
		{"a", true},
		{"hello", false},
		{"abccba", true},
		{"abcdcba", true},
	}

	for _, test := range tests {
		l := single.NewLinkList()

		for _, str := range test.sep {
			l.InsertTail(string(str))
		}
		out := IsPalindrome1(l)
		if out != test.want {
			if "" == test.sep {
				test.sep = "nil"
			}
			t.Errorf("err: IsPlalindrome1(LIST:%q) out=%t,want=%t", test.sep, out, test.want)
		}
	}
}


func TestIsPlalindrome2(t *testing.T) {
	tests := []struct {
		sep  string
		want bool
	}{
		{"", false},
		{"a", true},
		{"hello", false},
		{"abccba", true},
		{"abcdcba", true},
	}

	for _, test := range tests {
		l := single.NewLinkList()

		for _, str := range test.sep {
			l.InsertTail(string(str))
		}
		out := IsPalindrome2(l)
		if out != test.want {
			if "" == test.sep {
				test.sep = "nil"
			}
			t.Errorf("err: IsPlalindrome1(LIST:%q) out=%t,want=%t", test.sep, out, test.want)
		}
	}
}

func TestIsPlalindrome3(t *testing.T) {
	tests := []struct {
		sep  string
		want bool
	}{
		{"", false},
		{"a", true},
		{"hello", false},
		{"abccba", true},
		{"abcdcba", true},
	}

	for _, test := range tests {
		l := single.NewLinkList()

		for _, str := range test.sep {
			l.InsertTail(string(str))
		}
		out := IsPalindrome3(l)
		if out != test.want {
			if "" == test.sep {
				test.sep = "nil"
			}
			t.Errorf("err: IsPlalindrome1(LIST:%q) out=%t,want=%t", test.sep, out, test.want)
		}
	}
}