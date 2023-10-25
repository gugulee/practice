package gomonkey

type Computer struct {
}

func (t *Computer) NetworkCompute(a, b int) (int, error) {
	// do something in remote computer
	c := a + b

	return c, nil
}

func (t *Computer) Compute(a, b int) (int, error) {
	sum, err := t.NetworkCompute(a, b)
	return sum, err
}
