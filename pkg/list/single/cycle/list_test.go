package cycle

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestInsetHead(t *testing.T) {
	tests := []struct {
		sep  string
		want string
	}{
		{sep: "hello"},
	}
	for _, test := range tests {
		var want []string
		for _, s := range test.sep {
			want = append(want, string(s))
		}
		reverse(want)
		test.want = strings.Join(want, "->")

		l := NewLinkList()
		for _, str := range test.sep {
			l.InsertHead(string(str))
		}
		out := l.Print()
		if test.want != out {
			t.Errorf("err: LIST: %q,out=%q,want %q", test.sep, out, test.want)
		}
		t.Logf("info: LIST: %q,length=%d", test.sep, l.len)
	}
}

func TestInsetTail(t *testing.T) {
	tests := []struct {
		sep  string
		want string
	}{
		{sep: "hello"},
	}
	for _, test := range tests {
		var want []string
		for _, s := range test.sep {
			want = append(want, string(s))
		}
		test.want = strings.Join(want, "->")

		l := NewLinkList()
		for _, str := range test.sep {
			l.InsertTail(string(str))
		}
		out := l.Print()

		if test.want != out {
			t.Errorf("err: LIST: %q,out=%q,want %q", test.sep, out, test.want)
		}
		t.Logf("info: LIST: %q,length=%d", test.sep, l.len)
	}
}

func TestSearchListBynode(t *testing.T) {
	l := NewLinkList()
	strs := "hello"
	for _, str := range strs {
		l.InsertTail(string(str))
	}
	tests := []struct {
		sep string
	}{
		{sep: "h"},
		{sep: "o"},
		{sep: "a"},
		{sep: "c"},
	}
	for _, test := range tests {
		out := l.SearchListByValue(test.sep)
		if nil == out {
			if strings.Contains(strs, test.sep) {
				require.NotNilf(t, out, "LIST: %q, %q should be found", strs, test.sep)
				continue
			}
			t.Logf("LIST: %q,can not find %q", strs, test.sep)
		}
	}
}

func TestSearchListByValue(t *testing.T) {
	l := NewLinkList()
	node := l.SearchListByValue("h")
	require.Nil(t, node)

	strs := "hello"
	for _, str := range strs {
		l.InsertTail(string(str))
	}

	tests := []struct {
		sep string
	}{
		{sep: "h"},
		{sep: "o"},
		{sep: "a"},
		{sep: "c"},
	}

	for _, test := range tests {
		if ! strings.Contains(strs, test.sep) {
			continue
		}
		node := l.SearchListByValue(test.sep)
		require.NotNilf(t, node, "LIST: %q, %q should be found", strs, test.sep)
	}
}

func TestPrint(t *testing.T) {
	tests := []struct {
		sep  string
		want string
	}{
		{sep: "hello"},
		{sep: "world,lee"},
	}

	for _, test := range tests {
		var want []string
		for _, s := range test.sep {
			want = append(want, string(s))
		}
		test.want = strings.Join(want, "->")

		l := NewLinkList()
		for _, str := range test.sep {
			l.InsertTail(string(str))
		}
		out := l.Print()
		if test.want != out {
			t.Errorf("LIST: %q,out=%q,want %q", test.sep, out, test.want)
		}
	}
}
