package single

import (
	"testing"
)

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestLruCache(t *testing.T) {
	tests := []struct {
		sep  string
		want string
		len  uint32
	}{
		{"abcde", "e->d->c->b->a", 10},
		{"abcade", "e->d->a->c->b", 10},
		{"abcde", "e->d->c->b", 4},
		{"hello", "o->l->e->h", 5},
		{"hello", "o->l->e", 3},
		{"lhelo", "o->l->e", 3},
		{"lhelo", "o->l", 2},
		{"aaaaaabbbbbbbbb", "b->a", 10},
	}

	for _, test := range tests {
		lru := NewLru(test.len)
		for _, v := range test.sep {
			lru.LruCache(string(v))
		}
		out := lru.Print()
		if test.want != out {
			t.Errorf("err: LRU: %q,out %q,want %q", test.sep, out, test.want)
		} else {
			t.Logf("info: LRU: %q", out)
		}
	}

}
