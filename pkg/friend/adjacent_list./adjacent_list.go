package adjacentlist

import (
	"fmt"
	"strings"

	"github.com/practice/pkg/constants"
	skip "github.com/practice/pkg/list/skip/v3"
	vs "github.com/practice/pkg/utils/value_store"
)

const (
	defaultCapacity = 16
)

// Friend ...
type Friend struct {
	// follow store the users who follow
	follow []*skip.Skip

	// fans store the user's fans
	fans []*skip.Skip
}

// New ...
func New(capacity int) *Friend {
	if 0 == capacity {
		capacity = defaultCapacity
	}

	follow := make([]*skip.Skip, capacity)
	for i := range follow {
		follow[i] = skip.New(0, &vs.Ivalue{Value: constants.MinInt})
	}

	fans := make([]*skip.Skip, capacity)
	for i := range fans {
		fans[i] = skip.New(0, &vs.Ivalue{Value: constants.MinInt})
	}

	return &Friend{
		follow: follow,
		fans:   fans,
	}
}

// Insert the relation into table,
// and source follow target
func (f *Friend) Insert(source, target int) {
	f.follow[source].Insert(&vs.Ivalue{Value: target})

	f.fans[target].Insert(&vs.Ivalue{Value: source})
}

// Delete the relation from table,
// and source unfollow target
func (f *Friend) Delete(source, target int) {
	f.follow[source].Delete(&vs.Ivalue{Value: target})

	f.fans[target].Delete(&vs.Ivalue{Value: source})
}

// IsFollow return true if source follow target, else return false
func (f *Friend) IsFollow(source, target int) bool {
	return nil != f.follow[source].Search(&vs.Ivalue{Value: target})
}

// IsFans return true if source is target's fans, else return false
func (f *Friend) IsFans(source, target int) bool {
	return nil != f.fans[target].Search(&vs.Ivalue{Value: source})
}

// PrintFollow print the users who follow
func (f *Friend) PrintFollow() string {
	var r []string
	for i := range f.follow {
		s := f.follow[i].Head().PrintSlice(0)
		for j := range s {
			t := s[j].(*vs.Ivalue).String()
			r = append(r, fmt.Sprintf("%d-->%s", i, t))
		}
	}

	return strings.Join(r, ",")
}

// PrintFans print the user's fans
func (f *Friend) PrintFans() string {
	var r []string
	for i := range f.fans {
		s := f.fans[i].Head().PrintSlice(0)
		for j := range s {
			t := s[j].(*vs.Ivalue).String()
			r = append(r, fmt.Sprintf("%d-->%s", i, t))
		}
	}

	return strings.Join(r, ",")
}