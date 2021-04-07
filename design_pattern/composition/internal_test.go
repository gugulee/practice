package composition

import "testing"

func Test_test(t *testing.T) {
	ostrich := NewOstrich("ostrich")
	ostrich.tweet()
	ostrich.layEgg()

	sparrow := NewSparrow("sparrow")
	sparrow.tweet()
	sparrow.fly()
	sparrow.layEgg()
}
