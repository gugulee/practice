package composition

import "fmt"

type Flyable interface {
	fly()
}

type Tweetable interface {
	tweet()
}

type EggLayble interface {
	layEgg()
}

type Flyability struct {
	name string
}

func (f *Flyability) fly() {
	fmt.Printf("%s fly\n", f.name)
}

type TweetAbility struct {
	name string
}

func (t *TweetAbility) tweet() {
	fmt.Printf("%s tweet\n", t.name)
}

type EggLayAbility struct {
	name string
}

func (e *EggLayAbility) layEgg() {
	fmt.Printf("%s lay egg\n", e.name)
}

type Ostrich struct {
	name          string
	tweetAbility  *TweetAbility
	eggLayAbility *EggLayAbility
}

func NewOstrich(name string) *Ostrich {
	return &Ostrich{
		name:          name,
		tweetAbility:  &TweetAbility{name},
		eggLayAbility: &EggLayAbility{name},
	}
}

func (o *Ostrich) tweet() {
	o.tweetAbility.tweet()
}

func (o *Ostrich) layEgg() {
	o.eggLayAbility.layEgg()
}

type Sparrow struct {
	name          string
	flyability    *Flyability
	tweetAbility  *TweetAbility
	eggLayAbility *EggLayAbility
}

func NewSparrow(name string) *Sparrow {
	return &Sparrow{
		name:          name,
		flyability:    &Flyability{name},
		tweetAbility:  &TweetAbility{name},
		eggLayAbility: &EggLayAbility{name},
	}
}

func (s *Sparrow) fly() {
	s.flyability.fly()
}

func (s *Sparrow) tweet() {
	s.tweetAbility.tweet()
}

func (s *Sparrow) layEgg() {
	s.eggLayAbility.layEgg()
}
