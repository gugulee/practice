package template

import "testing"

func Test_Play(t *testing.T) {
	cricket := &Cricket{}
	play := New(cricket)
	play.Play()

	football := &Football{}
	play = New(football)
	play.Play()
}
