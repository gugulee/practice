package single

import (
	"testing"

	"github.com/practice/pkg/list/single"
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
			t.Errorf("err: IsPlalindrome2(LIST:%q) out=%t,want=%t", test.sep, out, test.want)
		}
	}
}

// todo:
func TestIsPlalindrome3(t *testing.T) {
	tests := []struct {
		sep  string
		want bool
	}{
		{"", false},
		{"a", true},
		{"healo", false},
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
			t.Errorf("err: IsPlalindrome3(LIST:%q) out=%t,want=%t", test.sep, out, test.want)
		}
	}
}
